package handlers

import (
	"github.com/Curicaveri/proxy-app/api/middleware"
	"github.com/kataras/iris"
)

// SetupRouter will handle the routes
func SetupRouter(app *iris.Application) {
	app.Get("/", middleware.ProxyMiddleware, proxyHandler)
}

func proxyHandler(c iris.Context) {
	c.JSON(iris.Map{"result": "ok"})
}
