package redis_client

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"net"
	"sogo/app/global/variable"
)

type LogHook struct{}

func (LogHook) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		return next(ctx, network, addr)
	}
}
func (LogHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		if err := cmd.Err(); err != nil {
			variable.RedisLog.Error("redis 日志", zap.String("cmd", err.Error()))
		} else {
			variable.RedisLog.Info("redis 日志", zap.String("cmd", cmd.String()))
		}
		err := next(ctx, cmd)
		if err != nil {
			return err
		}
		return nil
	}
}
func (LogHook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		return next(ctx, cmds)
	}
}
