package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		if err := c.HTML(http.StatusOK, "Uzaylılar tarafından kaçırıldım. Evet, tarafından."); err != nil {
			return err
		}
		return nil
	}).Name = "index"

	e.GET("/about", func(c echo.Context) error {
		if err := c.HTML(http.StatusOK, "Uzaylılar tarafından kaçırıldım. Evet, tarafından."); err != nil {
			return err
		}
		return nil
	})

	e.GET("/ping", func(c echo.Context) error {
		if err := c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"}); err != nil {
			return err
		}
		return nil
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	}).Name = "health"

	e.GET("/health/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	}).Name = "health-ping"

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
