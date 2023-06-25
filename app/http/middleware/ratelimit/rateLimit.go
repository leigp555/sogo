package ratelimit

import (
	"github.com/gin-gonic/gin"
	"sogo/app/global/consts"
	"sogo/app/global/variable"
	"sogo/app/utils/response"
	"time"
)

var (
	waitTime int
)

func init() {
	waitTime = variable.Config.GetInt("bucket.waitTime")

}

func RateLimitMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 没有令牌等待config.C.Bucket.WaitTime秒钟
		if !variable.Bucket.WaitMaxDuration(1, time.Second*time.Duration(waitTime)) {
			response.Fail(c, consts.CodeTooManyRequests, struct{}{})
			c.Abort()
			return
		}
		c.Next()
	}
}
