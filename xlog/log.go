package xlog

import "go.uber.org/zap"

var logger *zap.Logger

// logger 单例
func NewZapLogger() *zap.Logger {
	if logger == nil {
		logger = InitLogger()
	}
	return logger
}

func InitLogger() *zap.Logger {
	config := LogConfig{
		Filename:   "./runtime/app.log",
		MaxSize:    100, // 100 MB
		MaxBackups: 3,
		MaxAge:     7, // 7 days
		LocalTime:  true,
		Compress:   true,
		Level:      zap.DebugLevel,
		Encoding:   "json", // 可选 "console" 或留空使用默认格式
	}
	lg, err := Init(config)
	if err != nil {
		panic(err)
	}
	return lg
}

// info
func Info(msg ...any) {
	NewZapLogger().Sugar().Info(msg...)
}

func Infof(template string, args ...any) {
	NewZapLogger().Sugar().Infof(template, args...)
}

func Infow(msg string, keysAndValues ...any) {
	NewZapLogger().Sugar().Infow(msg, keysAndValues...)
}

// warn
func Warn(msg ...any) {
	NewZapLogger().Sugar().Warn(msg...)
}

func Warnf(template string, args ...any) {
	NewZapLogger().Sugar().Warnf(template, args...)
}

// debug
func Debug(msg ...any) {
	NewZapLogger().Sugar().Debug(msg...)
}

func Debugf(template string, args ...any) {
	NewZapLogger().Sugar().Debugf(template, args...)
}

// error
func Error(msg ...any) {
	NewZapLogger().Sugar().Error(msg...)
}

func Errorf(template string, args ...any) {
	NewZapLogger().Sugar().Errorf(template, args...)
}
