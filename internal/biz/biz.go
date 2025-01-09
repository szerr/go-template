package biz

import "github.com/google/wire"

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	wire.Struct(new(UserBiz), "*"),
)
