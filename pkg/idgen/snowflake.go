package idgen

import (
	"encoding/binary"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func init() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
}

// NewID returns a new snowflake ID as a 16-byte slice (binary(16) compatible).
func NewID() []byte {
	id := node.Generate().Int64()
	b := make([]byte, 16)
	binary.BigEndian.PutUint64(b[8:], uint64(id))
	return b
}
