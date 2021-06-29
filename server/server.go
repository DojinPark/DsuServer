package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/DojinPark/DuckServer/auth"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	e := echo.New()

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("./templates/*.html")),
	}
	e.Renderer = renderer

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"name": "Dolly!",
		})
	}).Name = "foobar"
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Welcome to scrollduck.com!")
	//})

	// Authentication Routes
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

	e.Logger.Fatal(e.Start(":1120"))
}
