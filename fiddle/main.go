package main

import (
	"net/http"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func hello(c *echo.Context) error {
	return c.String(http.StatusOK, "hello world\n")

}

func main() {
	//Echo instance
	e := echo.New()

	//middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.Gzip())

	//routes
	e.Get("/", hello)

	//open debug
	e.SetDebug(true)

	//start server
	e.Run(":1234")
}
