package main

import (
	"net/http"

	"MS_API/RESTAURANTS_MS/controller"
	"MS_API/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()
	storage.NewDB()

	e.GET("/restaurants", controller.GetRestaurants)
	// Start server
	e.Logger.Fatal(e.Start(":1324"))
}

func start(c echo.Context) error {
	return c.String(http.StatusOK, "Starting app........ ")
}
