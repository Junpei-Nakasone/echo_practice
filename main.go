package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from echo.\n")
}

func main() {
	fmt.Println("hello world from web server.")

	e := echo.New()

	e.GET("/", hello)
	e.Start(":8080")
}
