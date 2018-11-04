package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			getStock(),
		)
	})

	e.GET("/user/:id", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			getUserStock(c.Param("id")),
		)
	})

	e.POST("/user/:id/checkout", func(c echo.Context) error {
		checkoutBook(ISBN(c.QueryParam("isbn")), c.Param("id"))

		return c.NoContent(http.StatusOK)
	})

	e.POST("/user/:id/return", func(c echo.Context) error {
		returnBook(ISBN(c.QueryParam("isbn")), c.Param("id"))

		return c.NoContent(http.StatusOK)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
