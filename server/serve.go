package server

import "github.com/go-martini/martini"

func Runserver(port string) {
  //　Provides some reasonable defaults that work well for web application.
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
