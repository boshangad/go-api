package config

type App struct {
	// 根目录路径
	Root string
	// 接口地址
	ApiUrl string
	// 静态资源路径
	StaticUrl string
	// 文件上传路径
	UploadPath string
	// 用户token参数名称
	TokenParamName string
	// 应用参数名称
	AppParamName string
}
