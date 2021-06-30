package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterTemplate(e *echo.Echo) {
	e.GET("/", indexPage).Name = "/"
}

func indexPage(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "Dolly!",
	})
}
