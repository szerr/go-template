//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"go-template/internal/biz"
	"go-template/internal/data"
	"go-template/internal/pkg"
)

// The build tag makes sure the stub is not built in the final build.

// wireApp init
func wireApp() (*biz.UserBiz, func(), error) {
	panic(wire.Build(
		pkg.PkgProviderSet,
		data.DataProviderSet,
		biz.BizProviderSet,
	))
}
