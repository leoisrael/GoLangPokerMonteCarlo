package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/leoisrael/GoLangPokerMonteCarlo/prob"
)

func Api(port string) {
	e := echo.New()
	e.Static("/", "static")

	e.GET("/", func(c echo.Context) error {
		return c.File("static/index.html")
	})

	// Define a rota POST para o cálculo de probabilidade no flop
	e.POST("/api/probabilidade/flop", func(c echo.Context) error {
		entrada := new(prob.EntradaProbabilidade)
		if err := c.Bind(entrada); err != nil {
			return err
		}
		resultado := prob.ProbabilidadeFlop(*entrada)
		return c.String(http.StatusOK, resultado)
	})

	// Similarmente, defina as rotas para turn e river
	e.POST("/api/probabilidade/turn", func(c echo.Context) error {
		entrada := new(prob.EntradaProbabilidade)
		if err := c.Bind(entrada); err != nil {
			return err
		}
		resultado := prob.ProbabilidadeTurn(*entrada)
		return c.String(http.StatusOK, resultado)
	})

	e.POST("/api/probabilidade/river", func(c echo.Context) error {
		entrada := new(prob.EntradaProbabilidade)
		if err := c.Bind(entrada); err != nil {
			return err
		}
		resultado := prob.ProbabilidadeRiver(*entrada)
		return c.String(http.StatusOK, resultado)
	})

	e.Logger.Fatal(e.Start(":" + port))
}
