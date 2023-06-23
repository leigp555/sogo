package consts

type StatusCode int

// 自定义状态码
const (
	// CURD 常用业务状态码

	CurdStatusOkCode StatusCode = 200

	CurdCreatFailCode StatusCode = 1000 + iota
	CurdUpdateFailCode
	CurdDeleteFailCode
	CurdSelectFailCode
	CurdRegisterFailCode
	CurdLoginFailCode
	CurdRefreshTokenFailCode
	ServerOccurredErrorCode
	ValidatorParamsCheckFailCode

	CodeClientError StatusCode = 2000 + iota
	CodeUnauthorized
	CodeForbidden
	CodeNotFound
	CodeBusy
	CodeInvalidParam
	CodeUserNotExist
	CodeUsernameExist
	CodeEmailCaptchaErr
	CodeImgCaptchaErr
	CodeRequireToken
	CodeTokenFormatErr
	CodeInvalidToken
	CodeTooManyRequests
)

// 状态码描述
var codeMsgMap = map[StatusCode]string{
	CurdStatusOkCode:             "Success",
	CurdCreatFailCode:            "新增失败",
	CurdUpdateFailCode:           "更新失败",
	CurdDeleteFailCode:           "删除失败",
	CurdSelectFailCode:           "查询无数据",
	CurdRegisterFailCode:         "注册失败",
	CurdLoginFailCode:            "登录失败",
	CurdRefreshTokenFailCode:     "刷新Token失败",
	ServerOccurredErrorCode:      "服务器内部发生代码执行错误",
	ValidatorParamsCheckFailCode: "参数校验失败",
	CodeClientError:              "客户端错误",
	CodeUnauthorized:             "身份认证失败",
	CodeForbidden:                "无权限访问",
	CodeNotFound:                 "资源不存在",
	CodeBusy:                     "服务繁忙,请稍后再试",
	CodeInvalidParam:             "参数错误",
	CodeUserNotExist:             "用户名或密码不正确",
	CodeUsernameExist:            "用户名已存在",
	CodeEmailCaptchaErr:          "邮箱验证码错误",
	CodeImgCaptchaErr:            "图片验证码错误",
	CodeRequireToken:             "请上传登录凭证",
	CodeTokenFormatErr:           "token格式错误",
	CodeInvalidToken:             "无效的凭证",
	CodeTooManyRequests:          "请求频繁,请稍后再试",
}

func (c StatusCode) Msg() string {
	msg, exit := codeMsgMap[c]
	if !exit {
		msg = codeMsgMap[CodeBusy]
	}
	return msg
}
