package auth

import (
	"github.com/gin-gonic/gin"
	"sogo/app/global/consts"
	"sogo/app/utils/response"
	"sogo/app/utils/token"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//1.获取token
		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" {
			response.Fail(c, consts.CodeRequireToken, struct{}{})
			c.Abort()
			return
		}
		//2.分离出token数据
		splitArr := strings.Split(tokenHeader, " ")
		if len(splitArr) != 2 || splitArr[0] != "Bearer" {
			response.Fail(c, consts.CodeTokenFormatErr, struct{}{})
			c.Abort()
			return
		}
		tokenStr := splitArr[1]
		//3.解析token
		uid, err := token.ParseAccessToken(tokenStr)
		if err != nil {
			response.Fail(c, consts.CodeInvalidToken, struct{}{})
			c.Abort()
			return
		}
		//4.将uid放到上下文中方便后续使用
		c.Set(consts.UserId, uid)
		c.Next()
	}
}
