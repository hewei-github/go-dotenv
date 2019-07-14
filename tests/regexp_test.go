package tests

import (
	"github.com/hewei-github/go-dotenv"
	"os"
	"path/filepath"
	regexp2 "regexp"
	"strings"
	"testing"
	"time"
)

func TestRegexpVar(t *testing.T)  {
	var data = map[string]string{
		  "name" : "h",
		  "sex" : "man",
		  "age" : "100",
		  "app_id" : "${age}",
		  "app_name" :"${name}",
		  "app_root" :"${root}/app",
		  "root" : "app/",//CurrentDir(),
		  "storage_path" : "${app_root}/${app_name}/storage/",
		  "time" : time.Now().String(),
	}
	var regexp,err = regexp2.Compile(Env.DefVarExpress)
	if err !=nil {
		t.Error(err.Error())
	}
	for _,v :=range data  {
		if tmp:=regexp.FindAllStringSubmatch(v,-1); 0 != len(tmp) {
			t.Logf("match : %s , value: %s",tmp,v)
		}
	}
	if nil == t {

	}
}

func TestSplitN(t *testing.T)  {
	var strArr = [...]string{
		"var=123","name=str=1","base64=axf31sdf==",
		"abc","$={1232=123}","","123213",
	}
	for _,data :=range strArr {
		arr:=strings.SplitN(data,Env.DefAssignOpt, 2)
		t.Log(arr)
		Dump(arr)
	}
}

func CurrentDir() string  {
	if dir,err:=filepath.Abs(filepath.Dir(os.Args[0]));err == nil {
		return dir
	}
	return "."
}
