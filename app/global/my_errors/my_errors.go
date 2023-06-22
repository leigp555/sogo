package my_errors

const (
	ErrorStoreKeyAlreadyExist      string = "kvStore键名已经存在"
	ErrorsGormDBCreateParamsNotPtr string = "gorm Create 函数的参数必须是一个指针"
	ErrorsGormDBUpdateParamsNotPtr string = "gorm 的 Update、Save 函数的参数必须是一个指针"
	ErrorsGormInitFail             string = "Gorm 数据库驱动、连接初始化失败"
)
