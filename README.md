环境遍历读取包
--------------------
pakcage : hewei-github/go-dotenv

 ![license](https://img.shields.io/badge/license-MIT-yellow.svg)
 ![go](https://img.shields.io/badge/go-^1.11-green.svg)
 ![package](https://img.shields.io/badge/go-package-blue.svg)
 ![function](https://img.shields.io/badge/function-env-red.svg)
 ![test](https://img.shields.io/badge/go-test-green.svg)

> 支持.env 文件 .ini 文件读取 支持自动读取环境变量,支持配置动态变量(${xxx})

    .env 
     APP_NAME=test
     APP_ROOT=./
     APP_ENV=${APP_NAME}
    
    #代码
    eg : 
        
        import (
            "github.com/hewei-github/go-dotenv"
            "os"
         )
         
         var file :=".env"
         app :=Env.RootEnv.Load(file)
         if app == nil{
            // failed read env key value config
            return 
         }
         
         if "test"==app.Get("APP_NAME") {
             // ok
         }
         
         if app.IsSet("test") {
            // ok
         }
         
         if app.Get("os") == os.GetEnv("os") {
            // ok
         }
