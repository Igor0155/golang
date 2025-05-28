package main

import (
	"net/http"

	"github.com/Igor0155/golang/go_web/routes"
)

func main() {
	routes.HandleRequest()
	http.ListenAndServe(":8080", nil)
}
