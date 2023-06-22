package gorm_sql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sogo/app/global/variable"
)

func CreateMysqlClient() *gorm.DB {
	dial := getDbDial()
	db, err := gorm.Open(dial, &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 gormLog(), //拦截、接管 gorm v2 自带日志
	})
	if err != nil {
		panic(fmt.Sprintf("mysql connect error: %v", err))
	}
	return db
}

func getDbDial() gorm.Dialector {
	dns := getDns()
	dial := mysql.Open(dns)
	return dial
}

func getDns() (dns string) {
	username := variable.Config.GetString("mysql.username")
	password := variable.Config.GetString("mysql.password")
	host := variable.Config.GetString("mysql.host")
	port := variable.Config.GetInt("mysql.port")
	db := variable.Config.GetString("mysql.db")
	charset := variable.Config.GetString("mysql.charset")
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=false&loc=Local", username, password, host, port, db, charset)
}

func gormLog() logger.Interface {
	return customLog(SetInfoStrFormat("[info] %s\n"), SetWarnStrFormat("[warn] %s\n"), SetErrStrFormat("[error] %s\n"),
		SetTraceStrFormat("[traceStr] %s [%.3fms] [rows:%v] %s\n"), SetTraceWarnStrFormat("[traceWarn] %s %s [%.3fms] [rows:%v] %s\n"), SetTraceErrStrFormat("[traceErr] %s %s [%.3fms] [rows:%v] %s\n"))
}
