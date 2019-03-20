package internal

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type API struct {
	Echo          *echo.Echo
	storageClient IStorageClient
}

func NewAPI(s IStorageClient) *API {
	return &API{
		Echo:          echo.New(),
		storageClient: s,
	}
}

func (api *API) Configure() {
	api.Echo.GET("/", api.getInventory)
	api.Echo.GET("/user/:id", api.getUserInventory)
	api.Echo.POST("/user/:id/checkout", api.checkoutBook)
	api.Echo.POST("/user/:id/return", api.returnBook)
}

func (api *API) getInventory(c echo.Context) error {
	inv, err := GetInventory(api.storageClient)

	if err != nil {
		return fmt.Errorf("Failed to get inventory: %s", err)
	}

	return c.JSON(http.StatusOK, inv)
}

func (api *API) getUserInventory(c echo.Context) error {
	inv, err := GetUserInventory(
		api.storageClient,
		c.Param("id"),
	)

	if err != nil {
		return fmt.Errorf("Failed to get user inventory: %s", err)
	}

	return c.JSON(http.StatusOK, inv)
}

func (api *API) checkoutBook(c echo.Context) error {
	err := CheckoutBook(
		api.storageClient,
		c.QueryParam("isbn"),
		c.Param("id"),
	)

	if err != nil {
		return fmt.Errorf("Failed to checkout book: %s", err)
	}

	return c.NoContent(http.StatusOK)
}

func (api *API) returnBook(c echo.Context) error {
	err := ReturnBook(
		api.storageClient,
		c.QueryParam("isbn"),
		c.Param("id"),
	)

	if err != nil {
		return fmt.Errorf("Failed to return book: %s", err)
	}

	return c.NoContent(http.StatusOK)
}
