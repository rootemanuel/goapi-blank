package main

import (
	"github.com/kataras/iris"

	. "github.com/rootemanuel/goapi-blank/service"
)

var user = UserService{}

func main() {
	app := iris.Default()

	app.Get("/teste", user.Teste)

	app.Run(iris.Addr(":8080"))
}
