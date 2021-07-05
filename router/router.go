package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	m "github.com/DojinPark/DuckServer/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Renderer = m.Renderer()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//test:
	e.Use(middleware.CORS())

	return e
}
