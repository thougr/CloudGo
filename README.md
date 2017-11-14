# 使用Martini开发框架

---
### 使用Ｍartini的八大理由  
-  &nbsp;使用极其简单.  
-  &nbsp;无侵入式的设计.  
-  &nbsp;很好的与其他的Go语言包协同使用.  
-  &nbsp;超赞的路径匹配和路由.  
-  &nbsp;模块化的设计 - 容易插入功能件，也容易将其拔出来.  
-  &nbsp;已有很多的中间件可以直接使用.  
-  &nbsp;框架内已拥有很好的开箱即用的功能支持.  
-  &nbsp;完全兼容http.HandlerFunc接口.  

### 建立Ｍartini的环境

    go get github.com/go-martini/martini  
    
### 编程 web 服务程序  

    package server
    
    import "github.com/go-martini/martini"
    
    func Runserver(port string) {
      // Provides some reasonable defaults that work well for web application.
      m := martini.Classic()
      // Two handlers return something.
      m.Get("/", func() string {
        return "Hello world!"
      })
      m.Get("/hello/:name", func(params martini.Params) string {
      return "Hello " + params["name"]
    })
      // Run at port.
      m.RunOnAddr(":"+port)
    }  
  
  main函数使用老师提供的（作稍微的修改）。   

### 运行服务器
![此处输入图片的描述][1]

### 测试命令
![此处输入图片的描述][2]
此外接着还输出`Hello testuser`

### 压力测试
![此处输入图片的描述][3]


  [1]: https://raw.githubusercontent.com/thougr/CloudGo/master/screenshot/run.png
  [2]: https://raw.githubusercontent.com/thougr/CloudGo/master/screenshot/curltest.png
  [3]: https://raw.githubusercontent.com/thougr/CloudGo/master/screenshot/pressuretest.png
