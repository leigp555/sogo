package mysql_client

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"sogo/app/global/variable"
	"strings"
	"time"
)

type logOutPut struct{}

type gormLogger struct {
	logger.Writer
	logger.Config
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

func customLog(options ...Options) logger.Interface {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)
	logConf := logger.Config{
		SlowThreshold: time.Second * variable.Config.GetDuration("mysql.slowThreshold"),
		LogLevel:      logger.Info,
		Colorful:      true,
	}
	log := &gormLogger{
		Writer:       logOutPut{},
		Config:       logConf,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
	for _, val := range options {
		val.setConfig(log)
	}
	return log
}

// 添加可选参数

type Options interface {
	setConfig(*gormLogger)
}
type OptionFunc func(logConf *gormLogger)

// 这个函数作用是调用自己
func (f OptionFunc) setConfig(logConf *gormLogger) {
	f(logConf)
}

func SetInfoStrFormat(format string) Options {
	return OptionFunc(func(logConf *gormLogger) {
		logConf.infoStr = format
	})
}
func SetWarnStrFormat(format string) Options {
	return OptionFunc(func(logConf *gormLogger) {
		logConf.warnStr = format
	})
}
func SetErrStrFormat(format string) Options {
	return OptionFunc(func(logConf *gormLogger) {
		logConf.errStr = format
	})
}
func SetTraceStrFormat(format string) Options {
	return OptionFunc(func(logConf *gormLogger) {
		logConf.traceStr = format
	})
}
func SetTraceWarnStrFormat(format string) Options {
	return OptionFunc(func(logConf *gormLogger) {
		logConf.traceWarnStr = format
	})
}
func SetTraceErrStrFormat(format string) Options {
	return OptionFunc(func(logConf *gormLogger) {
		logConf.traceErrStr = format
	})
}
func (l logOutPut) Printf(strFormat string, args ...interface{}) {
	logRes := fmt.Sprintf(strFormat, args...)
	logFlag := "gorm 日志:"
	detailFlag := "详情："
	if strings.HasPrefix(strFormat, "[info]") || strings.HasPrefix(strFormat, "[traceStr]") {
		variable.MysqlLog.Info(logFlag, zap.String(detailFlag, logRes))
	} else if strings.HasPrefix(strFormat, "[error]") || strings.HasPrefix(strFormat, "[traceErr]") {
		variable.MysqlLog.Error(logFlag, zap.String(detailFlag, logRes))
	} else if strings.HasPrefix(strFormat, "[warn]") || strings.HasPrefix(strFormat, "[traceWarn]") {
		variable.MysqlLog.Warn(logFlag, zap.String(detailFlag, logRes))
	}
}

func (l *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info print info
func (l gormLogger) Info(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.Printf(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Warn print warn messages
func (l gormLogger) Warn(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.Printf(l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Error print error messages
func (l gormLogger) Error(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.Printf(l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Trace print sql message
func (l gormLogger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.LogLevel == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
