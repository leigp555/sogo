package bootstrap

import (
	"sogo/app/global/variable"
	"sogo/app/utils"
)

func InitDeps() {
	utils.InitLogger()
	variable.ZapLog.Info("Initializing dependencies")
}
