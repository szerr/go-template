package handle

import (
	"github.com/google/wire"
)

// @title go-template
// @version v0.0.1
// @description description
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http https
// @basePath /api/v1/

// ProviderSet is handle providers.
var HandleProviderSet = wire.NewSet(
	wire.Struct(new(AuthHandle), "*"),
	wire.Struct(new(UserHandle), "*"),
	wire.Struct(new(BaseHandle), "*"),
)
