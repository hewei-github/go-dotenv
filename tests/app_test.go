package tests

import (
	"fmt"
	"github.com/hewei-github/go-dotenv"
	"os"
	"reflect"
	"regexp"
	"testing"
)

func Assert(desc string, val bool, msg string, t *testing.T) {
	if !val {
		t.Errorf("%s tests not pass,error : %s", desc, msg)
	} else {
		t.Logf("%s tests passed", desc)
	}
}

func Log(msg string) {
	fmt.Println(msg)
}

func Dump(value interface{}) {
	Log(reflect.TypeOf(value).String())
}

func TestAppCreate(t *testing.T) {

	var file = Env.GetFile("./index.html")
	var flag = "go" == file.Extension()
	// Log(file.Extension())
	// Log(file.AbsolutePath())
	Dump(file.PathInfo())
	Log(file.FileType())
	Log(file.Extension())
	fmt.Println(file.PathInfo())
	Assert("文件类型测试", !flag, "文件类型获取异常", t)
}

func TestMapGet(t *testing.T) {
	var a = make(map[string]interface{})
	a["name"] = "hello"
	a["test"] = 1
	var key = "name"
	if val, ok := a[key]; ok {
		Assert("测试map获取值", val == "hello", "测试失败", t)
	}
}

func TestGet(t *testing.T)  {
	var app=Env.RootEnv.Load("dev.env")
	if app == nil{
		t.Errorf("load env file failed")
		return
	}
	var testTables = map[string]interface{}{
		"os" : os.Getenv("os"),
		"time" : "1562507355",
		"JAVA_HOME":os.Getenv("JAVA_HOME"),
		"name" : "h",
		"null":"",
	}
	for k,v:=range testTables {
		if val:=app.Get(k,nil); val!= v {
			t.Errorf("get env key[%s] expect value:%s ,but real value:%s",k,v,val)
		}
	}
}

//func TestGetEnvApp(t *testing.T) {
//	var extArr = []string{"./.env", "./index.html", "./app.ini", "./1.txt", "../App.go", "test.yaml"}
//	for _, it := range extArr {
//		if _, err := Env.GetEnvApp(it); nil != err {
//			t.Error(err.Error())
//		}
//	}
//}

func TestEnvGetAndSet(t *testing.T) {
	app, _ := Env.GetEnvApp(".env")
	data := map[string]interface{}{
		"test": "test",
		"arr":  []int{12, 3, 4, 7},
		"map":  map[string]string{"name": "123", "arr1": "[1,2,3,4]"},
	}
	for k, v := range data {
		app.Set(k, v)
	}
	for k, v := range data {
		if Env.ToString(v) != Env.ToString(app.Get(k, nil)) {
			t.Errorf("set key value : %s=>%s failed ", k, v)
		} else {
			fmt.Printf("%s=>%s \n", k, v)
		}
	}
}

func getTestData() map[string]string {
	var data = map[string]string{
		"name":         "h",
		"sex":          "man",
		"age":          "${os}",
		"app_id":       "${age}",
		"app_name":     "${name}",
		"app_root":     "${root}/app",
		"root":       "./app/", //   CurrentDir(),
		"storage_path": "${app_root}/${app_name}/storage/",
		"time":       "1562507355" , //Env.ToString(time.Now().Unix()),
	}
	return data
}

func TestEnvSave(t *testing.T) {
	app, _ := Env.GetEnvApp(".env")
	data := getTestData()
	for k, v := range data {
		app.Set(k, v)
	}
	if 0 != app.Save("dev.env", os.ModeAppend) {
		t.Log("env save pass")
	} else {
		t.Error("env save failed")
	}
}

func TestEnvLoad(t *testing.T) {
	app := Env.RootEnv.Load("dev.env")
	data := getTestData()
	if app == nil {
		t.Error("load env file failed")
	}
	reg, _ := regexp.Compile(Env.DefVarExpress)
	// Dump(app.GetAll())
	for k, v := range data {
		// fmt.Printf("%s=>%s \n",k,v)
		if !app.IsSet(k) {
			t.Logf("key [ %s ] fail to set (value:%s)", k, v)
			continue
		}
		if val := Env.ToString(app.Get(k, nil)); val != v && !reg.MatchString(v) {
			t.Logf("key [ %s ] value error , expect : %s ,but real get value : %s", k, v, val)
		}
	}

	// for k, v := range app.GetAll() {
	// 	fmt.Printf("%s=%s \n", k, v)
	// }
}

func BenchmarkGetEnvApp(b *testing.B) {

	var extArr = []string{"./.env", "./dev.env"}
	b.Run("create-object", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, it := range extArr {
				if app := Env.RootEnv.Load(it); nil != app {
				}
			}
		}
	})
}
