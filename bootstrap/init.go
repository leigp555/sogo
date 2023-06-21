package bootstrap

import (
	"sogo/app/global/variable"
	"sogo/app/utils/zap_factory"
)

func InitDeps() {
	variable.GinLog = zap_factory.CreateLogger("gin")
	variable.ZapLog = zap_factory.CreateLogger("system")
	variable.MysqlLog = zap_factory.CreateLogger("mysql")
}
