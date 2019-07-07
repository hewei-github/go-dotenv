package Env

import (
	"fmt"
	"reflect"
	"strings"
)

// 是否对应类型
func IsType(value interface{},ty string)  bool {
	if t:=GetType(value); t== strings.ToLower(ty) || t==ty {
		return true
	}
	return false
}

// 获取类型
func GetType(value interface{}) string  {
	t :=fmt.Sprintf("%s",reflect.TypeOf(value))
	if t == "%!s(<nil>)" {
		return "nil"
	}
	return t
}