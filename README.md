环境遍历读取包
--------------------
hewei-github/go-dotenv

> 支持.env 文件 .ini 文件读取
    
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
