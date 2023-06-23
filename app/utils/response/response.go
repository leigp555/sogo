package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sogo/app/global/consts"
)

// ReturnJson 通用响应
func ReturnJson(C *gin.Context, httpCode int, dataCode consts.StatusCode, msg string, data interface{}) {
	id := C.GetString(consts.RequestId)
	C.Header("Content-Type", "application/json; charset=utf-8")
	C.JSON(httpCode, gin.H{
		"requestId": id,
		"code":      dataCode,
		"msg":       msg,
		"data":      data,
	})
}

// 语法糖函数封装

// Success(c,CurdStatusOkCode,"example")

// Success 直接返回成功
func Success(c *gin.Context, data interface{}) {
	msgCode := consts.CurdStatusOkCode
	msg := msgCode.Msg()
	ReturnJson(c, http.StatusOK, consts.CurdStatusOkCode, msg, data)
}

// Fail(c,CurdStatusOkCode,"example")

// Fail 失败的业务逻辑
func Fail(c *gin.Context, msgCode consts.StatusCode, data interface{}) {
	msg := msgCode.Msg()
	ReturnJson(c, http.StatusBadRequest, msgCode, msg, data)
	c.Abort()
}
