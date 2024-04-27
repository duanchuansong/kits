package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisConfig 存储Redis连接所需的配置信息
type RedisConfig struct {
	Address     string        // Redis服务器地址，如 "localhost:6379"
	Password    string        // 密码（可选）
	DB          int           // 数据库编号（可选，默认为0）
	PoolSize    int           // 连接池大小（可选，默认为10）
	MinIdle     int           // 最小空闲连接数（可选，默认为0）
	MaxIdle     int           // 最大空闲连接数（可选，默认为不设置）
	IdleTimeout time.Duration // 空闲连接超时时间（可选，默认为5分钟）

	DialTimeout  time.Duration // 建立连接超时时间（可选，默认为5秒）
	ReadTimeout  time.Duration // 读取数据超时时间（可选，默认为3秒）
	WriteTimeout time.Duration // 写入数据超时时间（可选，默认为3秒）
}

// ConnectRedis 封装连接Redis的函数
func ConnectRedis(cfg *RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         cfg.Address,
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdle,
		MaxIdleConns: cfg.MaxIdle,

		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		//IdleTimeout:  cfg.IdleTimeout,
	})

	// 检查连接是否正常
	ctx, cancel := context.WithTimeout(context.Background(), cfg.DialTimeout)
	defer cancel()

	err := rdb.Ping(ctx).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to ping Redis: %w", err)
	}

	log.Println("Connected to Redis successfully")
	return rdb, nil
}
