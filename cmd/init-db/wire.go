//go:build wireinject
// +build wireinject

package main

import (
	"github.com/casbin/casbin/v2"
	"github.com/google/wire"
	"go-template/internal/pkg"
	"go-template/internal/pkg/permissions"
)

// The build tag makes sure the stub is not built in the final build.

// wireApp init
func wireApp() (*casbin.Enforcer, func(), error) {
	panic(wire.Build(
		pkg.PkgProviderSet,
		permissions.NewCasbinEnforcer,
	))
}
