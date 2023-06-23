package redis_client

import (
	"context"
	"github.com/redis/go-redis/v9"
	"sogo/app/global/my_errors"
	"sogo/app/global/variable"
	"time"
)

func NewRedisClient() *redis.Client {
	addr := variable.Config.GetString("redis.addr")
	password := variable.Config.GetString("redis.password")
	db := variable.Config.GetInt("redis.db")
	poolSize := variable.Config.GetInt("redis.poolSize")

	rdb := redis.NewClient(&redis.Options{
		Addr:                  addr,
		Password:              password,
		DB:                    db, // use default DB
		MaxRetries:            3,
		MinRetryBackoff:       8 * time.Millisecond,
		MaxRetryBackoff:       512 * time.Millisecond,
		DialTimeout:           5 * time.Second,
		ReadTimeout:           3 * time.Second,
		WriteTimeout:          3 * time.Second,
		ContextTimeoutEnabled: false,
		PoolFIFO:              false,
		PoolSize:              poolSize, // 连接池大小
		MinIdleConns:          5,
		MaxIdleConns:          10,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(my_errors.ErrorsRedisInitFail + err.Error())
	}

	rdb.AddHook(LogHook{})

	return rdb
}
