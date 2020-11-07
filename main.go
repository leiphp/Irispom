package main

import (
	"Irispom/cron"
	"github.com/kataras/iris/v12"
	"sync"
	"time"
)

func main() {
	//生产mq,每秒去生产一个消息
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 1; i<=10;i++  {
		go cron.InitProduce(i, &wg)
		time.Sleep(1*time.Second)
	}
	wg.Wait()

	//消费mq
	//go cron.InitConsumer()

	app := iris.Default()
	app.Use(myMiddleware)

	app.RegisterView(iris.HTML("./views", ".html"))

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})

	app.Get("/", func(ctx iris.Context) {
		// Bind: {{.message}} with "Hello world!"
		ctx.ViewData("message", "Hello world!")
		// Render template file: ./views/hello.html
		ctx.View("hello.html")
	})

	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
		userID, _ := ctx.Params().GetUint64("id")
		ctx.Writef("User ID: %d", userID)
	})

	// Listens and serves incoming http requests
	// on http://localhost:8080.
	app.Run(iris.Addr(":8181"))
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}