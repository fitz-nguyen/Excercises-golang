package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	// Cache certificates
	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
		<h1>Welcome to Echo!</h1>
		<h3>TLS certificates automatically installed from Let's Encrypt :)</h3>
	`)
	})

	// start server
	e.Logger.Fatal(e.Start(":1323"))

}
