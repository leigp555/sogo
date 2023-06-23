package bootstrap

import (
	"sogo/app/global/variable"
	"sogo/app/utils/elasticsearch_client"
	"sogo/app/utils/mysql_client"
	"sogo/app/utils/redis_client"
	"sogo/app/utils/snow_flake"
	"sogo/app/utils/zap_factory"
)

func InitDeps() {
	//初始化日志
	variable.GinLog = zap_factory.CreateLogger("gin")
	variable.ZapLog = zap_factory.CreateLogger("system")
	variable.MysqlLog = zap_factory.CreateLogger("mysql")
	variable.RedisLog = zap_factory.CreateLogger("redis")

	//初始化mysql
	variable.Mdb = mysql_client.CreateMysqlClient()

	//初始化redis
	variable.Rdb = redis_client.NewRedisClient()

	//雪花算法节点
	variable.SnowNode = snow_flake.NewSnowflake()

	//初始化elasticsearch
	variable.Es = elasticsearch_client.NewElasticsearchClient()
}
