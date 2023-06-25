package ginLogger

import "C"
import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sogo/app/global/consts"
	"sogo/app/global/variable"
	"time"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//开始时间
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		if variable.Config.GetString("system.env") == "dev" {
			variable.GinLog.Info("\033[35m"+path+"\033[0m",
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Duration("cost", cost),
			)
		} else {
			variable.GinLog.Info(path,
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("referer", c.Request.Referer()),
				zap.String("request_id", c.GetString(consts.RequestId)),
				zap.Int64("requestSize", c.Request.ContentLength),
				zap.Int("responseSize", c.Writer.Size()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Duration("cost", cost),
			)
		}

	}
}
