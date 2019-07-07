package Env

// 环境变量
type ObjectEnv struct {
	file File
	container map[string]interface{}
	parser ParserService
}

// 默认解析器
type DefaultParser struct {}

// 解析器
type ParserService interface {
	GetContent(file string,assignOpt string,varExpress string) (map[string]interface{},error)
}

// 环境接口
type IEnvironment interface {
	Get(key string,def interface{}) interface{}
	Set(key string, value interface{}) * ObjectEnv
	Load(path string) * ObjectEnv
}

// 根
var RootEnv * ObjectEnv

func init()  {
	RootEnv = new(ObjectEnv)
	RootEnv.parser = DefaultParser{}
}