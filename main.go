package main

import (
	"github.com/hspearman/library-inventory/internal"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	api := echo.New()
	storageClient := internal.NewStorageClient()

	internal.SetDefaultInventory(storageClient)

	api.GET("/", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			internal.GetInventory(storageClient),
		)
	})

	api.GET("/user/:id", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			internal.GetUserInventory(
				storageClient,
				c.Param("id"),
			),
		)
	})

	api.POST("/user/:id/checkout", func(c echo.Context) error {
		internal.CheckoutBook(
			storageClient,
			c.QueryParam("isbn"),
			c.Param("id"),
		)

		return c.NoContent(http.StatusOK)
	})

	api.POST("/user/:id/return", func(c echo.Context) error {
		internal.ReturnBook(
			storageClient,
			c.QueryParam("isbn"),
			c.Param("id"),
		)

		return c.NoContent(http.StatusOK)
	})

	api.Logger.Fatal(api.Start(":1323"))
}
