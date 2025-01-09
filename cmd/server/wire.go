//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"go-template/api/engine"
	"go-template/api/handle"
	v1 "go-template/api/routers/v1"
	"go-template/internal/biz"
	"go-template/internal/data"
	"go-template/internal/pkg"
	"go-template/internal/pkg/permissions"
)

// The build tag makes sure the stub is not built in the final build.

// wireApp init
func wireApp() (func() error, func(), error) {
	panic(wire.Build(
		pkg.PkgProviderSet,
		data.DataProviderSet,
		permissions.NewCasbinEnforcer,
		biz.BizProviderSet,
		handle.HandleProviderSet,
		v1.RouterPkgProviderSet,
		engine.EngineProviderSet,
		newApp,
	))
}
