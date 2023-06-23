package validator_trans

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"sogo/app/global/consts"
	"sogo/app/global/my_errors"
	"sogo/app/utils/response"
	"strings"
)

// ValidatorError 翻译表单参数验证器出现的校验错误
func ValidatorError(c *gin.Context, err error) {
	if errs, ok := err.(validator.ValidationErrors); ok {
		wrongParam := removeTopStruct(errs.Translate(trans))
		response.Fail(c, consts.ValidatorParamsCheckFailCode, wrongParam)
	} else {
		errStr := err.Error()
		// multipart:nextpart:eof 错误表示验证器需要一些参数，但是调用者没有提交任何参数
		if strings.ReplaceAll(strings.ToLower(errStr), " ", "") == "multipart:nextpart:eof" {
			response.Fail(c, consts.ValidatorParamsCheckFailCode, map[string]string{"tips": my_errors.ErrorNotAllParamsIsBlank})
		} else {
			response.Fail(c, consts.ValidatorParamsCheckFailCode, map[string]string{"tips": errStr})
		}
	}
}
