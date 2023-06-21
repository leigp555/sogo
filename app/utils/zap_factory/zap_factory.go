package zap_factory

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sogo/app/global/variable"
	"time"
)

type options struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	ShowLine   bool
	Level      string
}

func createOption(k string) *options {
	return &options{
		Filename:   variable.Config.GetString(k + ".logConfig.filename"),
		MaxSize:    variable.Config.GetInt(k + ".logConfig.maxSize"),
		MaxBackups: variable.Config.GetInt(k + ".logConfig.maxBackups"),
		MaxAge:     variable.Config.GetInt(k + ".logConfig.maxAge"),
		Compress:   variable.Config.GetBool(k + ".logConfig.compress"),
		ShowLine:   variable.Config.GetBool(k + ".logConfig.showLine"),
		Level:      variable.Config.GetString(k + ".logConfig.level"),
	}
}

func CreateLogger(k string) (log *zap.Logger) {
	opt := createOption(k)
	return generateLog(opt)
}

// 根据配置生成zap日志对象

func generateLog(opt *options) *zap.Logger {
	writeSyncer := getLogWriter(opt.Filename, opt.MaxSize, opt.MaxBackups, opt.MaxAge, opt.Compress)
	encoder := getEncoder()
	level := getLogLevel(opt.Level)
	core := zapcore.NewCore(encoder, writeSyncer, level)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段,如：添加一个服务器名称
	//filed := zap.Fields(zap.String("user", logConf.Prefix))
	// 构造日志
	var logger *zap.Logger
	if opt.ShowLine {
		logger = zap.New(core, caller, development)
	} else {
		logger = zap.New(core, caller)
	}
	return logger
}

func getEncoder() zapcore.Encoder {
	var encodeLevel zapcore.LevelEncoder
	if variable.Config.GetString("system.env") == "dev" {
		encodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		encodeLevel = zapcore.LowercaseLevelEncoder
	}
	//自定义时间格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	//自定义代码路径、行号输出
	customCallerEncoder := func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + caller.TrimmedPath() + "]")
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     "\n",
		EncodeLevel:    encodeLevel,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   customCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	if variable.Config.GetString("system.env") == "dev" {
		return zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		return zapcore.NewJSONEncoder(encoderConfig)
	}

}
func getLogWriter(f string, ms int, mb int, ma int, cp bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   f,
		MaxSize:    ms,
		MaxBackups: mb,
		MaxAge:     ma,
		Compress:   cp,
	}
	if variable.Config.GetString("system.env") == "dev" {
		//return  zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger)) //既输出到文件又输出到控制台
		return zapcore.AddSync(os.Stderr) //开发模式下输出到控制台
	} else {
		return zapcore.AddSync(lumberJackLogger) //生产环境下输出到文件
	}

}
func getLogLevel(l string) zapcore.Level {
	var level zapcore.Level
	switch l {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	return level
}
