package pkg

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/v2/util/grand"
)

var Snowflake, _ = snowflake.NewNode(int64(grand.N(1, 1023)))
