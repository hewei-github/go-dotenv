package tests

import (
	Env "github.com/hewei-github/go-dotenv"
	"testing"
)

func TestToString(t *testing.T)  {
	var str string
	var hashMap = make(map[string]interface{})
	hashMap["number"] = 123
	hashMap["name"] = "name"
	hashMap["interface"] = new(interface{})
	hashMap["func"] = func(v string) string { return v+"string" }
	hashMap["bool"] = true
	hashMap["object"] = Env.GetFile(".env")
	hashMap["中文"] = "你好好看看"
	hashMap["null"] = nil
	for key,val:=range hashMap  {
		str=Env.ToString(val)
		t.Logf("%s=>%s",key,str)
	}
	if fun,ok:=hashMap["func"].(func(v string) string); ok {
		 t.Logf(fun("test-->>>"))
	}
}