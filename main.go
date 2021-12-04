package main

import (
	"capstone/app/api/routes"
	"capstone/app/config"
	"capstone/app/database"

	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	database.Init(e)
	routes.InitializeRoutes(e)

	port := config.GetConfig(e).Port
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}