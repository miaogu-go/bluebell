package snowflake

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func Init(startTime string, machineId int64) (err error) {
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineId)

	return
}

func GetID() int64 {
	return node.Generate().Int64()
}
