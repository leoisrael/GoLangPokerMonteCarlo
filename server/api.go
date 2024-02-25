package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Api(port string) {
	e := echo.New()
	e.Static("/", "static")

	e.GET("/", func(c echo.Context) error {
		return c.File("static/index.html")
	})

	e.GET("/api/poker", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "API de Poker com Echo!",
		})
	})

	e.Logger.Fatal(e.Start(":" + port))
}
