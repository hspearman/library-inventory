package main

import (
	"fmt"
	"github.com/hspearman/library-inventory/internal"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func main() {
	api := echo.New()

	storageClient, err := internal.NewStorageClient()
	if err != nil {
		log.Fatalf("Failed to dial redis: %s", err.Error())
	}

	defer storageClient.Close()

	internal.SetDefaultInventory(storageClient)

	api.GET("/", func(c echo.Context) error {
		inv, err := internal.GetInventory(storageClient)

		if err != nil {
			return fmt.Errorf("Failed to get inventory: %s", err)
		}

		return c.JSON(http.StatusOK, inv)
	})

	api.GET("/user/:id", func(c echo.Context) error {
		inv, err := internal.GetUserInventory(
			storageClient,
			c.Param("id"),
		)

		if err != nil {
			return fmt.Errorf("Failed to get user inventory: %s", err)
		}

		return c.JSON(http.StatusOK, inv)
	})

	api.POST("/user/:id/checkout", func(c echo.Context) error {
		err := internal.CheckoutBook(
			storageClient,
			c.QueryParam("isbn"),
			c.Param("id"),
		)

		if err != nil {
			return fmt.Errorf("Failed to checkout book: %s", err)
		}

		return c.NoContent(http.StatusOK)
	})

	api.POST("/user/:id/return", func(c echo.Context) error {
		err := internal.ReturnBook(
			storageClient,
			c.QueryParam("isbn"),
			c.Param("id"),
		)

		if err != nil {
			return fmt.Errorf("Failed to return book: %s", err)
		}

		return c.NoContent(http.StatusOK)
	})

	api.Logger.Fatal(api.Start(":1323"))
}
