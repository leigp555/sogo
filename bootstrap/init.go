package bootstrap

import (
	"sogo/app/global/variable"
	"sogo/app/utils/gorm_sql"
	"sogo/app/utils/zap_factory"
)

func InitDeps() {
	//初始化日志
	variable.GinLog = zap_factory.CreateLogger("gin")
	variable.ZapLog = zap_factory.CreateLogger("system")
	variable.MysqlLog = zap_factory.CreateLogger("mysql")

	//初始化mysql
	variable.Mdb = gorm_sql.CreateMysqlClient()
}
