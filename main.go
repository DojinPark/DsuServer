package main

import (
	"github.com/DojinPark/DuckServer/handler"
	"github.com/DojinPark/DuckServer/middleware"
	"github.com/DojinPark/DuckServer/router"
)

func main() {
	e := router.New()
	handler.RegisterTemplate(e)

	g := e.Group("/auth")
	middleware.RegisterAuth(g)

	e.Logger.Fatal(e.Start(":1120"))
}
