package variable

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sogo/app/global/types"
)

var (
	// Config 全局配置
	Config types.YamlConf

	//全局日志
	ZapLog   *zap.Logger
	GinLog   *zap.Logger
	MysqlLog *zap.Logger

	//全局Store前缀
	ConfigKeyPrefix = "Config_"

	//全局mysql数据库
	Mdb *gorm.DB
)
