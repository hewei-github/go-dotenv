package Env

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// 获取
func (app *ObjectEnv) Get(key string, def interface{}) interface{} {
	if val, ok := app.container[key]; ok {
		return val
	}
	if val := os.Getenv(key); "" != val {
		return val
	}
	return def
}

// 配置键个数
func (app *ObjectEnv) Size() int {
	if app != nil {
		return len(app.container)
	}
	return 0
}

// 设置
func (app *ObjectEnv) Set(key string, value interface{}) *ObjectEnv {
	app.container[key] = value
	return app
}

// 加载
func (app ObjectEnv) Load(path string) *ObjectEnv {
	var file *File
	file = GetFile(path)
	if nil == file || !file.IsExist() {
		return nil
	}
	if ok, err := file.IsDir(); ok || err != nil {
		return nil
	}
	env := new(ObjectEnv)
	env.parser = RootEnv.parser
	if content, err := app.parser.GetContent(path, DefAssignOpt, DefVarExpress); nil == err {
		env.container = content
	}
	return env
}

// 删除键
func (app *ObjectEnv) Unset(key string) bool {
	if app.IsSet(key) {
		delete(app.container, key)
		return true
	}
	return false
}

// 是否存在
func (app *ObjectEnv) IsSet(key string) bool {
	if app.container == nil {
		return false
	}
	if _, ok := app.container[key]; ok {
		return true
	}
	return false
}

// 保存
func (app *ObjectEnv) Save(file string, mode os.FileMode) int {
	var content string
	for key, value := range app.container {
		line := key + DefAssignOpt + ToString(value) + LineDiv
		if "" == content {
			content = line
		} else {
			content = content + line
		}
	}
	var size = len(content)
	if err := ioutil.WriteFile(file, []byte(content), mode); nil == err {
		return size
	}
	return 0
}

// 获取所有
func (app *ObjectEnv) GetAll() map[string]string {
	var data = make(map[string]string)
	if 0 == len(app.container) {
		return data
	}
	for k, v := range app.container {
		data[k] = ToString(v)
	}
	return data
}

// 创建
func GetEnvApp(path string) (*ObjectEnv, error) {
	file := GetFile(path)
	if !file.IsExist() {
		return nil, fmt.Errorf("file: " + path + " not exists")
	}
	if !RootEnv.IsSupport(file.path) {
		return nil, fmt.Errorf("file: " + path + " not support for env parser")
	}
	return RootEnv.Load(file.AbsolutePath()), nil
}

// 是否支持
func (app *ObjectEnv) IsSupport(path string) bool {
	if path == "" {
		return false
	}
	file := GetFile(path)
	if val, err := file.IsDir(); err != nil || val {
		return false
	}
	var ext = file.Extension()
	for _, value := range SupportExtArr {
		if value == ext {
			return true
		}
	}
	return false
}

// 解析器
func (_ DefaultParser) GetContent(file string, assignOpt string, varExpress string) (map[string]interface{}, error) {
	content := make(map[string]interface{})
	// println("get content "+file)
	if data, err := FileGetContent(file); nil == err {
		if 0 == len(data) {
			err := fmt.Errorf("%s content is empty", file)
			return content, err
		}
		reg, err := regexp.Compile(varExpress)
		if nil != err {
			reg = nil
		}
		if dataMap, err := ParseStrLine(data, assignOpt); nil == err {
			for k, v := range dataMap {
				if v != "" && reg != nil {
					v = value(v, reg, dataMap)
				}
				// println("set data",v)
				content[k] = v
			}
		} else {
			return content, err
		}
	}
	return content, nil
}

// 特殊字符串变量解析
func value(str string, reg *regexp.Regexp, context map[string]string) string {
	if str == "" || reg == nil {
		return str
	}
	if !reg.MatchString(str) {
		return str
	}
	if tmp := reg.FindAllStringSubmatch(str, -1); 0 != len(tmp) {
		for _, arr := range tmp {
			if 3 != len(arr) {
				continue
			}
			key := arr[2]
			rep := arr[1]
			if key == "" {
				continue
			}
			if val, ok := context[key]; ok {
				if reg.MatchString(val) {
					val = value(val, reg, context)
				}
				str = strings.Replace(str, rep, val, -1)
				continue
			}
			if val := os.Getenv(key); val != "" {
				if reg.MatchString(val) {
					val = value(val, reg, context)
				}
				str = strings.Replace(str, rep, val, -1)
			}
		}
	}
	return str
}

// 解析表达式字符串
func ParseKeyValueByExpressStr(data string, assignOpt string) (map[string]string, bool) {
	if data == "" || assignOpt == "" {
		return nil, false
	}
	result := strings.SplitN(data, assignOpt, 2)
	if 2 == len(result) {
		return map[string]string{
			result[0]: result[1],
		}, true
	}
	return nil, false
}

// 换行拆解
func ParseStrLine(data string, assignOpt string) (map[string]string, error) {
	var err error
	var content map[string]string
	lines := strings.Split(data, LineDiv)
	if 0 == len(lines) {
		err = fmt.Errorf("%s content ,parse failed ", "no match")
		return nil, err
	}
	content = make(map[string]string)
	for _, v := range lines {
		if item, ret := ParseKeyValueByExpressStr(v, assignOpt); !ret {
			continue
		} else {
			for key, value := range item {
				content[key] = value
			}
		}
	}
	if 0 != len(content) {
		return content, nil
	}
	err = fmt.Errorf("%s ,parse failed", "key value express not match")
	return nil, err
}

// 包加载初始化函数
func init()  {
	RootEnv = new(ObjectEnv)
	RootEnv.parser = DefaultParser{}
}