package mysql

import (
	"context"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// WithZapLogger 使用zap日志记录到文件
func WithZapLogger(logPath string, level zapcore.Level) Option {
	return func(cfg *gorm.Config) {
		// 构建zap logger
		writeSyncer := getZapLogWriter(logPath) // 移除了多余的下划线和逗号
		zapLogger := zap.New(zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			writeSyncer,
			level,
		))

		// 将zap logger适配为GORM的Logger
		cfg.Logger = NewZapGormLogger(zapLogger)
	}
}

// getZapLogWriter 设置zap日志文件写入器
func getZapLogWriter(logPath string) zapcore.WriteSyncer {
	rotateWriter, err := rotatelogs.New(
		logPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(logPath),          // 生成软链接
		rotatelogs.WithMaxAge(24*time.Hour*7),     // 保留7天日志
		rotatelogs.WithRotationTime(24*time.Hour), // 每天轮转一次
	)
	if err != nil {
		panic(err)
	}
	return zapcore.AddSync(rotateWriter)
}

// NewZapGormLogger 创建一个zap的日志适配器，符合gorm的Logger接口
func NewZapGormLogger(zapLogger *zap.Logger) logger.Interface {
	return &zapGormLogger{zapLogger: zapLogger}
}

type zapGormLogger struct {
	zapLogger *zap.Logger
}

func (l *zapGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *zapGormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.zapLogger.Sugar().Infof(msg, data...)
}

func (l *zapGormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.zapLogger.Sugar().Warnf(msg, data...)
}

func (l *zapGormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.zapLogger.Sugar().Errorf(msg, data...)
}

func (l *zapGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	// 实现Trace方法，根据需要记录SQL执行跟踪信息
}
