package my_errors

const (
	ErrorStoreKeyAlreadyExist string = "kvStore键名已经存在"

	ErrorsGormDBCreateParamsNotPtr string = "gorm Create 函数的参数必须是一个指针"
	ErrorsGormDBUpdateParamsNotPtr string = "gorm 的 Update、Save 函数的参数必须是一个指针"
	ErrorsGormInitFail             string = "Gorm 数据库驱动、连接初始化失败"

	ErrorSnowFlakeInitFail string = "SnowFlake 初始化失败"

	ErrorsRedisInitFail string = "Redis 连接初始化失败"

	ErrorsElasticSearchInitFail string = "Redis 连接初始化失败"

	ErrorNotAllParamsIsBlank string = "该接口不允许所有参数都为空,请按照接口要求提交必填参数"

	ErrorsFuncEventAlreadyExists string = "注册函数类事件失败，键名已经被注册"
	ErrorsFuncEventNotRegister   string = "没有找到键名对应的函数"
	ErrorsFuncEventNotCall       string = "注册的函数无法正确执行"
	ErrorsBasePath               string = "初始化项目根目录失败"
	ErrorsTokenBaseInfo          string = "token最基本的格式错误,请提供一个有效的token!"
	ErrorsNoAuthorization        string = "token鉴权未通过，请通过token授权接口重新获取token,"
	ErrorsRefreshTokenFail       string = "token不符合刷新条件,请通过登陆接口重新获取token!"
	ErrorsParseTokenFail         string = "解析token失败"
	ErrorsCasbinNoAuthorization  string = "Casbin 鉴权未通过，请在后台检查 casbin 设置参数"

	ErrorsValidatorTransInitFail string = "validator的翻译器初始化错误"
)
