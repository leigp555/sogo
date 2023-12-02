package bootstrap

import (
	"log"
	"sogo/app/global/my_errors"
	"sogo/app/global/variable"
	"sogo/app/utils/bucket"
	"sogo/app/utils/elasticsearch_client"
	"sogo/app/utils/mysql_client"
	"sogo/app/utils/redis_client"
	"sogo/app/utils/snow_flake"
	"sogo/app/utils/validator_trans"
	"sogo/app/utils/zap_factory"
)

func InitDeps() {
	//初始化日志
	variable.GinLog = zap_factory.CreateLogger("gin")
	variable.ZapLog = zap_factory.CreateLogger("system")
	variable.MysqlLog = zap_factory.CreateLogger("mysql")
	variable.RedisLog = zap_factory.CreateLogger("redis")

	//初始化mysql客户端
	variable.Mdb = mysql_client.CreateMysqlClient()

	//初始化redis客户端
	variable.Rdb = redis_client.NewRedisClient()

	//雪花算法节点
	variable.SnowNode = snow_flake.NewSnowflake()

	//初始化elasticsearch客户端
	variable.Es = elasticsearch_client.NewElasticsearchClient()

	//全局注册 validator 错误翻译器,zh 代表中文，en 代表英语
	err := validator_trans.InitTrans("zh")
	if err != nil {
		log.Fatal(my_errors.ErrorsValidatorTransInitFail + err.Error())
	}

	//初始化一个限流桶
	variable.Bucket = bucket.NewBucket()
}
