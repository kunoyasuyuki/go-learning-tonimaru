package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type myHandler struct{}

// ServeHTTP implements the http.Handler interface.
func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ServeHTTP")
}

// main is the entry point for the application.
func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World from Echo!")
	})

	e.GET("/users", func(c echo.Context) error {
		return c.String(200, "Users!")
	})

	e.Start(":8080")

	//fmt.Println("Hello, World!")
	//
	//handler := myHandler{}
	//http.Handle("/", handler)
	//http.ListenAndServe(":8080", nil)
}
