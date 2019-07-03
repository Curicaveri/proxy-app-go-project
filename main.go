package main

import (
	"github.com/Curicaveri/proxy-app/api/handlers"
	"github.com/Curicaveri/proxy-app/api/middleware"
	"github.com/Curicaveri/proxy-app/api/server"
	"github.com/Curicaveri/proxy-app/api/utils"
)

func main() {
	utils.LoadEnv()
	app := server.SetUp()
	handlers.SetupRouter(app)
	middleware.InitQueue()
	server.RunServer(app)
}
