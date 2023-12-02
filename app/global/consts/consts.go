package consts

// 终端颜色控制
//const (
//	Reset  = "\033[0m"
//	Red    = "\033[31m"
//	Green  = "\033[32m"
//	Yellow = "\033[33m"
//	Blue   = "\033[34m"
//	Purple = "\033[35m"
//	Cyan   = "\033[36m"
//	Gray   = "\033[37m"
//)

var (
	ServerOccurredErrorMsg string = "服务器内部发生代码执行错误, "
	GinSetTrustProxyError  string = "Gin 设置信任代理服务器出错"

	RequestId string = "requestId"
	Purple           = "\033[35m"
	Reset            = "\033[0m"
	UserId           = "userId"
)
