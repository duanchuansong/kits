package mysql

import (
	"errors"
	"fmt"
	"github.com/duanchuansong/kits/xlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// MySQLConfig 存储MySQL连接所需的配置信息
type MySQLConfig struct {
	Host        string
	Port        int
	User        string
	Password    string
	Database    string
	MaxIdleConn int
	MaxOpenConn int
	// 这里可以添加更多的配置项
}

// Option 定义用于配置连接的选项函数类型
type Option func(*gorm.Config)

// WithLoggerLevel 设置日志级别
func WithLoggerLevel(level logger.LogLevel) Option {
	return func(cfg *gorm.Config) {
		cfg.Logger = logger.Default.LogMode(level)
	}
}

// WithPrepareStmt 是否开启预编译语句
func WithPrepareStmt(enable bool) Option {
	return func(cfg *gorm.Config) {
		cfg.PrepareStmt = enable
	}
}

// ConnectMySQL 封装GORM连接MySQL的函数，支持选项配置
func ConnectMySQL(cfg *MySQLConfig, opts ...Option) (*gorm.DB, error) {
	// 验证配置
	if cfg.User == "" || cfg.Password == "" || cfg.Host == "" || cfg.Database == "" {
		return nil, errors.New("MySQL configuration is invalid, make sure all fields are populated")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	// 注意：在生产环境中，建议使用环境变量或配置文件管理敏感信息（如用户名和密码）
	gormCfg := &gorm.Config{
		PrepareStmt: true,                                // 默认开启预编译
		Logger:      logger.Default.LogMode(logger.Info), // 默认日志级别为Info
	}
	for _, opt := range opts {
		opt(gormCfg)
	}

	db, err := gorm.Open(mysql.Open(dsn), gormCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB instance: %w", err)
	}
	// 设置连接池配置
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)

	// 连接测试
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping MySQL: %w", err)
	}

	// 在此记录成功连接的日志，在实际应用中可能需要更复杂的日志记录策略
	xlog.Info("Successfully connected to MySQL.")

	return db, nil
}
