package v1

import "github.com/google/wire"

// PkgProviderSet is routers providers.
var RouterPkgProviderSet = wire.NewSet(
	wire.Struct(new(UserRouter), "*"),
)
