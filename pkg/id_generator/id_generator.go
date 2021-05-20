package id_generator

import (
	"github.com/bwmarrin/snowflake"
	"nku-treehole-server/pkg/logger"
)

var node *snowflake.Node

func init() {
	n, err := snowflake.NewNode(1)
	if err != nil {
		logger.Fatalf("cannot init id_generator %v", err)
	}
	node = n
}

func GenerateID() int64 {
	id := node.Generate()
	return id.Int64()
}
