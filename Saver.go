package Env

import (
	"io/ioutil"
	"os"
)

// 写文件
func Write(file string, data string) int {
	return FilePutContent(file, data)
}

// 读文件
func Read(file string, buf interface{}) (data []byte, err error) {
	if content, err := FileGetContent(file); nil == err {
		if nil != buf {
			buf = content
		}
		return []byte(content), nil
	} else {
		return nil, err
	}
}

// 读取文件
func FileGetContent(file string) (content string, err error) {
	if content, err := ioutil.ReadFile(file); nil != err {
		return "", err
	} else {
		return string(content), nil
	}
}

// 写文件
func FilePutContent(file string, data string) int {
	var size = len(data)
	if "" == file || 0 == size {
		return 0
	}
	if err := ioutil.WriteFile(file, []byte(data), os.ModeAppend); nil != err {
		return 0
	}
	return size
}
