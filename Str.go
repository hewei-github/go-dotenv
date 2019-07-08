package Env

import (
	"fmt"
	"reflect"
	"strconv"
)

// 字符串 类型
type StrInterface interface {
	ToString(value interface{}) string
	StrToAny(value string) interface{}
}

// 字符串
type Str string

// 其他类型转字符串
func (_ Str) ToString(value interface{}) string {
	return ToString(value)
}

// 字符串转其他类型
func (_ Str) StrToAny(value string) interface{} {
	return StrToAny(value)
}

// 转字符串
func ToString(value interface{}) string {

	switch value.(type) {
		case string:
			return value.(string)
		case int:
			return strconv.Itoa(value.(int))
		case int8:
			return string(value.(int8))
		case int16:
			return string(value.(int16))
		case int32:
			return string(value.(int16))
		case int64:
			return strconv.FormatInt(value.(int64),10)
		case uint:
			return string(value.(uint))
		case uint8:
			return string(value.(uint8))
		case uint16:
			return string(value.(uint16))
		case uint32:
			return string(value.(uint32))
		case uint64:
			return strconv.FormatUint(value.(uint64),10)
		case bool:
			return strconv.FormatBool(value.(bool))
		case uintptr:
			return fmt.Sprintf("%x",value.(uintptr))
		case complex64:
			return fmt.Sprintf("%v",value.(complex64))
		case complex128:
			return fmt.Sprintf("%v",value.(complex128))
		case interface{}:
			return fmt.Sprintf("%s",value)
		case nil:
			return "nil"
		case func():
			return fmt.Sprintf("%s",reflect.ValueOf(value))
		case struct{}:
			return fmt.Sprintf("%s",reflect.ValueOf(value))
		default:
			return fmt.Sprintf("%s",reflect.ValueOf(value))
	}
}

// @todo 待完善
func StrToAny(value string) interface{} {
	if "" == value {
		return nil
	}
	return new(interface{})
}
