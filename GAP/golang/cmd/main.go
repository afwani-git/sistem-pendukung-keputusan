package main

import (
	"net/http"

	"github.com/afwani-git/sistem-pendukung-keputusan/GAP/golang/pkg/gap/factory"
	"github.com/labstack/echo/v4"
)

func main() {

	result := factory.LoadDummy()

	// echo
	e := echo.New()
	e.GET("/init", func(c echo.Context) error {
		return c.JSON(http.StatusOK, result.GapSpec)
	})

	e.GET("/gap-mapping", func(c echo.Context) error {

		return c.JSON(http.StatusOK, result.GapMapping)
	})

	e.GET("/calculate", func(c echo.Context) error {
		return c.JSON(http.StatusOK, result.GapCalculate)
	})

	e.GET("/ranking", func(c echo.Context) error {
		return c.JSON(http.StatusOK, result.GapRanking)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
