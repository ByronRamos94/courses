package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// instantiate echo
	e := echo.New()

	// route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "HELLO WORLD")

	})

	e.Logger.Fatal(e.Start(":1323"))
}
