package main

import (
	"github.com/DojinPark/DuckServer/handler"
	"github.com/DojinPark/DuckServer/middleware"
	"github.com/DojinPark/DuckServer/router"
)

func main() {
	e := router.New()

	handler.RegisterTemplate(e)
	middleware.RegisterAuth(e)

	e.Logger.Fatal(e.Start(":1120"))
}
