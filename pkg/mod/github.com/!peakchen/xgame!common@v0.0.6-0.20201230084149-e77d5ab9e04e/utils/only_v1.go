package utils

import (
	"github.com/bwmarrin/snowflake"
)

func NewOnly_v1() (onlyid snowflake.ID, err error) {
	var node *snowflake.Node
	node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
		return
	}

	// Generate a snowflake ID.
	onlyid = node.Generate()
	err = nil
	return
}

func NewInt64_v1() int64 {
	id, _ := NewOnly_v1()
	return id.Int64()
}

func NewString_v1() string {
	id, _ := NewOnly_v1()
	return id.String()
}
