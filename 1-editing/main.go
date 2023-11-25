package main

import (
	"fmt"
	"net/http"
)

type myHandler struct{}

// ServeHTTP implements the http.Handler interface.
func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ServeHTTP")
}

// main is the entry point for the application.
func main() {
	fmt.Println("Hello, World!")

	handler := myHandler{}
	http.Handle("/", handler)
	http.ListenAndServe(":8080", nil)
}
