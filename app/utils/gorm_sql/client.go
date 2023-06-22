package gorm_sql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sogo/app/global/my_errors"
	"sogo/app/global/variable"
	"time"
)

func CreateMysqlClient() *gorm.DB {
	dial := getDbDial()
	db, err := gorm.Open(dial, &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 gormLog(), //拦截、接管 gorm v2 自带日志
	})
	if err != nil {
		panic(my_errors.ErrorsGormInitFail + err.Error())
	}

	// 查询没有数据，屏蔽 gorm v2 包中会爆出的错误
	// https://github.com/go-gorm/gorm/issues/3789  此 issue 所反映的问题就是我们本次解决掉的
	_ = db.Callback().Query().Before("gorm:query").Register("disable_raise_record_not_found", MaskNotDataError)

	// https://github.com/go-gorm/gorm/issues/4838
	_ = db.Callback().Create().Before("gorm:before_create").Register("CreateBeforeHook", CreateBeforeHook)
	// 为了完美支持gorm的一系列回调函数
	_ = db.Callback().Update().Before("gorm:before_update").Register("UpdateBeforeHook", UpdateBeforeHook)

	rawDb, err := db.DB()
	if err != nil {
		panic(my_errors.ErrorsGormInitFail + err.Error())
	}

	// 连接池
	rawDb.SetConnMaxIdleTime(time.Second * 30)
	rawDb.SetConnMaxLifetime(variable.Config.GetDuration("mysql.setConnMaxLifetime") * time.Second)
	rawDb.SetMaxIdleConns(variable.Config.GetInt("mysqql.setMaxIdleConns"))
	rawDb.SetMaxOpenConns(variable.Config.GetInt("mysql.setMaxOpenConns"))

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
