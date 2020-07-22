package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "hello world",
		})
	})
	app.Get("/hello", func(ctx iris.Context) {
		r := ctx.GetReferrer()

		ctx.JSON(iris.Map{
			"teshello": "hello test",
			"r":        r,
		})

	})
	app.Run(iris.Addr(":8084"), iris.WithoutServerError(iris.ErrServerClosed))
}
