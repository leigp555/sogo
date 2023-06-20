package variable

import (
	"go.uber.org/zap"
	"sogo/app/global/types"
)

var (
	// Config 全局配置
	Config   types.YamlConf
	ZapLog   *zap.Logger
	GinLog   *zap.Logger
	MysqlLog *zap.Logger

	ConfigKeyPrefix = "Config_"
)
