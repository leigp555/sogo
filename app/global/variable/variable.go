package variable

import (
	"github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sogo/app/global/types"
)

var (
	DateFormat = "2006-01-02 15:04:05" //  设置全局日期时间格式

	// Config 全局配置
	Config types.YamlConf

	//全局日志
	ZapLog   *zap.Logger
	GinLog   *zap.Logger
	MysqlLog *zap.Logger
	RedisLog *zap.Logger

	//全局Store前缀
	ConfigKeyPrefix = "Config_"

	//全局mysql数据库
	Mdb *gorm.DB
	Rdb *redis.Client

	//雪花算法
	SnowNode *snowflake.Node
)
