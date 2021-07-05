package middleware

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterAuth(e *echo.Echo) {
	e.POST("/login", login)
	e.GET("/logout", logout)
	e.GET("/restorelogin", restoreLogin)

	g := e.Group("/restricted")
	config := middleware.JWTConfig{
		Claims:     &JWTCustomClaims{},
		SigningKey: []byte("my-secret"),
	}
	g.Use(middleware.JWTWithConfig(config))
	g.GET("", restricted)
}

type JWTCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	//test: Mock username, password
	//todo: apply MySQL DB
	if username != "test" || password != "test" {
		return echo.ErrUnauthorized
	}

	//test: Mock User Info
	//todo: apply MySQL DB
	//todo: set proper ExpiresAt
	claims := &JWTCustomClaims{
		"김지수",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	encoded, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": encoded,
	})
}

func logout(c echo.Context) error {
	//todo: define logout processes
	return c.JSON(http.StatusOK, echo.Map{
		"message": "logged out",
	})
}

func restoreLogin(c echo.Context) error {
	//todo: implement token comparison
	//return
}

func restricted(c echo.Context) error {
	//test: Mock accessible message
	//todo: send auth token, instead
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTCustomClaims)
	name := claims.Name

	return c.String(http.StatusOK, "User's Name: "+name)
}
