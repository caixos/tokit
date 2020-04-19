package constants

const (
	Err string = "9999::错误"

 	ErrStop string = "9990::暂停服务"
	ErrMQTTConnect string = "9991::MQTT链接失败"
	ErrCacheInit string = "9992::Cache初始化错误"
	ErrLoadEnv string = "9993::ENV环境加载失败"
	ErrSign string = "9994::数据签名错误"
	ErrBalance = "9995::金额不可为负值"

	ErrNotExist string = "9000::数据不存在"
	ErrIsExist string = "9001::数据已存在"
	ErrSaveFailed string = "9002::数据保存失败"

	ErrConvert string = "9010::类型转换错误"
	ErrJson string = "9011::JSON报文解析错误"

	ErrNoToken string = "9020::缺少authToken"
	ErrTokenFmt string = "9021::Token格式错误"
	ErrTokenExp string = "9022::Token过期,请重新登录"
	ErrTokenSign string = "9023::签名错误,请重新登录"

	ErrRoute string = "9030::路由解析错误"

)