//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"go-template/internal/model"
	"go-template/internal/pkg"
)

// The build tag makes sure the stub is not built in the final build.

// wireApp init
func wireApp() (func() error, func(), error) {
	panic(wire.Build(
		pkg.PkgProviderSet,
		model.ModelProviderSet,
		newApp,
	))
}
