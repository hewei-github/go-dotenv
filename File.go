package Env

import (
	"fmt"
	"mime"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

// 文件
type File struct {
	path string
	base map[string]string
}

// 文件协议
type FileSchema interface {
	IsExist() bool
	PathInfo() map[string]string
	IsDir() (bool, error)
	BaseName() string
	AbsolutePath() string
	FileType() string
	Extension() string
}

// 初始化对象
var FileNil = File{}

// 文件是否存在
func (file *File) IsExist() bool {
	if file.path == "" {
		return false
	}
	_, err := os.Stat(file.path)
	if err == nil {
		return true
	}
	if os.IsExist(err) {
		return false
	}
	return false
}

// path-info
func (file *File) PathInfo() map[string]string {
	if 0 == len(file.base) {
		file.inits()
	}
	return file.base
}

// 创建
func GetFile(path string) *File {
	var file = &File{
		path: path,
		base: make(map[string]string),
	}
	file = file.inits()
	return file
}

// 是否文件目录
func (file *File) IsDir() (bool, error) {
	if nil == file {
		return false,fmt.Errorf("nil pointer to call File IsDir")
	}
	if fs, err := os.Stat(file.path); nil == err {
		return fs.IsDir(), nil
	} else {
		return false, err
	}
}

// 初始化
func (file *File) inits() *File {
	if fs, err := os.Stat(file.path); nil == err {
		isDir := fs.IsDir()
		if isDir {
			file.base["isDir"] = "true"
			file.base["extension"] = ""
		} else {
			file.base["isDir"] = "false"
			file.base["extension"] = strings.Replace(path.Ext(file.path), ".", "", 1)
		}
		file.base["filename"] = fs.Name()
		file.base["size"] = strconv.FormatInt(fs.Size(), 10)
		file.base["mtime"] = fs.ModTime().String()
		file.base["mode"] = fs.Mode().String()
		file.base["basename"] = path.Base(file.path)
		file.base["dirname"] = path.Dir(file.path)
	}
	return file
}

// 获取基础文件信息
func (file *File) BaseName() string {
	if ok, err := file.base["basename"]; !err {
		file.inits()
	} else {
		return string(ok)
	}
	if ok, err := file.base["basename"]; err {
		return string(ok)
	}
	return ""
}

// 获取绝对路径
func (file *File) AbsolutePath() string {
	if file.path != "" && file.IsExist() {
		if abs, err := filepath.Abs(file.path); err == nil {
			return abs
		}
	}
	return ""
}

// 文件类型
func (file *File) FileType() string {
	if file.IsExist() {
		ext := path.Ext(file.path)
		if ext == "" {
			if typeName, mimeMap, err := mime.ParseMediaType(file.path); nil == err {
				if name, err := mimeMap[typeName]; err {
					return name
				}
			}
		}
		if ok, err := file.IsDir(); err == nil {
			if ok {
				return "dir"
			}
		}
		name := mime.TypeByExtension(ext)
		if "" != name {
			return name
		}
	}
	return ""
}

// 文件扩展类型
func (file *File) Extension() string {
	info := file.PathInfo()
	if value, ok := info["extension"]; ok {
		return value
	}
	return ""
}
