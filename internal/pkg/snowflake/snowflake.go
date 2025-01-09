package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"go-template/internal/pkg/config"
)

var node *snowflake.Node

func Init(conf *config.Config) error {
	var err error
	node, err = snowflake.NewNode(conf.App.NodeId)
	return err
}

func Int64() int64 {
	return node.Generate().Int64()
}
