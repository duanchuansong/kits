package xlog

import (
	"path/filepath"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogConfig 包含日志配置信息
type LogConfig struct {
	Filename   string        // 日志文件路径（不含日期部分）
	MaxSize    int           // 单个日志文件最大大小（MB）
	MaxBackups int           // 最大备份文件数量
	MaxAge     int           // 文件最多保留天数
	LocalTime  bool          // 是否使用本地时间
	Compress   bool          // 是否压缩旧日志文件
	Level      zapcore.Level // 日志级别
	Encoding   string        // 输出格式（json、console或zap's default）
}

// Init 初始化带有日志滚动功能（按天轮转）的Zap Logger
func Init(config LogConfig) (*zap.Logger, error) {
	// 计算当前日期对应的日志文件名
	today := time.Now().Format("20060102")
	fullFilename := filepath.Join(filepath.Dir(config.Filename), today+"."+filepath.Base(config.Filename))

	// lumberjack日志滚动配置
	logWriter := &lumberjack.Logger{
		Filename:   fullFilename,
		MaxSize:    config.MaxSize, // MB
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge, // days
		LocalTime:  config.LocalTime,
		Compress:   config.Compress,
	}

	// zap Encoder配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var encoder zapcore.Encoder
	switch config.Encoding {
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	case "console":
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// zap Core配置
	core := zapcore.NewCore(encoder, zapcore.AddSync(logWriter), config.Level)

	// 构建Logger
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return logger, nil
}
