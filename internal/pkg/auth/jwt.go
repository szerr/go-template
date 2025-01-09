package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go-template/internal/pkg/config"
	"go-template/internal/pkg/er"
	"go.uber.org/zap"
	"time"
)

type UserClaims struct {
	UserId uint64 `json:"user_id"`
	jwt.RegisteredClaims
}

type IJWT interface {
	GenToken(uint64) (string, error)
	ParseToken(string) (*UserClaims, error)
	RefreshToken(string) (string, error)
}

func NewJWT(c *config.Config) IJWT {
	return &JWT{conf: &c.Jwt}
}

type JWT struct {
	conf *config.Jwt
}

func (a *JWT) genJwt(claims *UserClaims) (string, error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(a.conf.AccessExpire) * time.Second)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    a.conf.Issuer,
		ID:        uuid.NewString(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	ss, err := token.SignedString([]byte(a.conf.Secret))
	return ss, er.WSEF(err, zap.String("secret", a.conf.Secret))
}

func (a *JWT) GenToken(userId uint64) (string, error) {
	return a.genJwt(&UserClaims{
		UserId: userId,
	})
}

func (a *JWT) ParseToken(t string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(t, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.conf.Secret), nil
	})

	if err != nil {
		return nil, er.WSEF(err, zap.Any("token", t))
	}

	if claims, ok := token.Claims.(*UserClaims); ok {
		return claims, nil
	} else {
		return nil, er.InvalidToken.WSF(zap.Any("claims", claims))
	}
}

func (a *JWT) RefreshToken(t string) (string, error) {
	claims, err := a.ParseToken(t)
	if err != nil {
		return "", er.WSEF(err, zap.String("token", t))
	}
	return a.genJwt(claims)
}
