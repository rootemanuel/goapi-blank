package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"

	. "github.com/rootemanuel/goapi-blank/config"
	. "github.com/rootemanuel/goapi-blank/dao"
)

var config = Config{}
var dao = RootDAO{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	app := iris.Default()

	app.Get("/teste", func(ctx context.Context) {

		rtn, err := dao.FindAll()
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		}

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(rtn)
	})

	app.Run(iris.Addr(":8080"))
}
