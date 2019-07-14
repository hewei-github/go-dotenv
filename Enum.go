package Env

// 包常量
const DefAssignOpt = "=" // 键值分隔符
const LineDiv ="\r\n" // 换行分隔符
const DefVarExpress = "(\\$\\{([^(\\{\\})]+)\\})" // 动态模版变量匹配
var SupportExtArr = [2]string{"env","ini"} // 支持文件后缀
var RootEnv * ObjectEnv // 引用根