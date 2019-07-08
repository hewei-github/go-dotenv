package tests

import (
	"github.com/hewei-github/go-dotenv"
	"strings"
	"testing"
)

func TestFileGetContent(t *testing.T) {
	var file = "dev.env"
	if content, err := Env.FileGetContent(file); nil == err {
		var arr []string
		/*go func(arr * []string) {*/
		lines := strings.Split(content, Env.LineDiv)
		arr = lines
		/*}(&arr)*/
		for _, v := range arr {
			println(v)
		}
	} else {
		t.Errorf("get file %s content fail", file)
	}
}

func TestFilePutContent(t *testing.T) {
	var file = "1.txt"
	var content = `
			[test]
			APP_NAME=env
			APP_ROOT=${project_dir}
			APP_KEY=1212121
	`
	if 0 != Env.FilePutContent(file, content) {
		t.Log("write content ok")
	} else {
		t.Error("write file failed")
	}
}
