package snow_flake

import (
	"github.com/bwmarrin/snowflake"
	"sogo/app/global/my_errors"
	"sogo/app/global/variable"
)

func NewSnowflake() *snowflake.Node {
	machineId := variable.Config.GetInt64("snowFlake.snowFlakeMachineId")
	node, err := snowflake.NewNode(machineId)
	if err != nil {
		panic(my_errors.ErrorSnowFlakeInitFail + err.Error())
	}
	return node
}
