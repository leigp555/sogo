package requestTrace

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"sogo/app/global/consts"
)

func RequestTrace() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestId := uuid.New().String()
		context.Set(consts.RequestId, requestId)
	}
}
