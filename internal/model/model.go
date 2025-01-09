package model

import (
	"github.com/google/wire"
)

// ModelProviderSet is model providers.
var ModelProviderSet = wire.NewSet(
	NewModel,
)

type AllModel []any

func NewModel() AllModel {
	return AllModel{
		new(SysUser),
	}
}
