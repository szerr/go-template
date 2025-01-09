package data

import (
	"github.com/google/wire"
	"go-template/internal/biz"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	NewBaseRepo,
	NewUserRepo,
	wire.Bind(new(biz.IUserRepo), new(*UserRepo)),
)
