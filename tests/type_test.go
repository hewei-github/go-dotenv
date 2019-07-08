package tests

import (
	"github.com/hewei-github/go-dotenv"
	"testing"
)

func getTestCases() map[string]interface{}  {
	cases :=make(map[string]interface{})
	cases["func-not-args"] = func() {}
	cases["func-has-args"] = func(v string) {}
	cases["func-has-args-with-return"] = func(v int) bool{ return v> 0  }
	cases["string-case"] = "testing-string"
	cases["int-case"] = 1
	cases["bool-case"] = true
	cases["nil-case"] = nil
	cases["struct-case"] = struct {}{}
	cases["object-case"] = Env.GetFile("a.ini")
	cases["interface-case"] = new(interface{})
	return cases
}

func TestGetType(t *testing.T)  {
	 tests := getTestCases()
	for desc,value:=range tests {
		ty := Env.GetType(value)
		if "" == ty {
			t.Errorf("%s fail ",desc)
		}else{
			t.Log(desc,ty)
		}
	}
}


