package variable

import (
	"github.com/bwmarrin/snowflake"
	es "github.com/elastic/go-elasticsearch/v7"
	"github.com/juju/ratelimit"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DateFormat = "2006-01-02 15:04:05" //  设置全局日期时间格式

	// Config 全局配置
	Config *viper.Viper

	//全局日志
	ZapLog   *zap.Logger
	GinLog   *zap.Logger
	MysqlLog *zap.Logger
	RedisLog *zap.Logger

	//全局Store前缀
	ConfigKeyPrefix = "Config_"

	//全局数据库
	Mdb *gorm.DB
	Rdb *redis.Client

	Es *es.Client

	//雪花算法
	SnowNode *snowflake.Node

	Bucket *ratelimit.Bucket
)
