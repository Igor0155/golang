package routes

import (
	"net/http"

	"github.com/Igor0155/golang/go_web/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)

}
