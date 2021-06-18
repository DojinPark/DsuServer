package main

import (
	"net/http"

	"github.com/DojinPark/DuckServer/auth"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Authentication Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome")
	})
	e.POST("/login", auth.Login)
	//	e.POST("/signup", Signup)
	//	e.GET("/logout", Logout)
	//	e.GET("/forgot", Forgot)

	// Mock auth area
	r := e.Group("/authaccess")
	config := middleware.JWTConfig{
		Claims:     &auth.JWTCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", auth.AuthAccess)

	e.Logger.Fatal(e.Start(":1323"))
}
