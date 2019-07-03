package server

import (
	"os"

	"github.com/kataras/iris"
)

// SetUp will configure the Iris Web Server
func SetUp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")
	return app
}

// RunServer will run the web application
func RunServer(app *iris.Application) {
	app.Run(
		iris.Addr(os.Getenv("PORT")),
	)
}
