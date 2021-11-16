package main

import (
	"DB_In_Memory/routes"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := getEnv("PORT", ":3000")
	App := echo.New()
	App.Use(middleware.Logger())
	App.Use(middleware.Recover())

	routes.Router(App)
	App.Logger.Fatal(App.Start(port))
}

func getEnv(key, defaultValue string) string {

	value, defined := os.LookupEnv(key)
	if !defined {
		return defaultValue
	}

	return value
}
