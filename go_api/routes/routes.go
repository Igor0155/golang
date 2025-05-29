package routes

import (
	"log"
	"net/http"

	"github.com/Igor0155/golang/go_api/controllers"
	"github.com/Igor0155/golang/go_api/middleware"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.PersonalidadePorId).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.CriarPersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletarPersonalidade).Methods("DELETE")
	r.HandleFunc("/api/personalidades/{id}", controllers.AtualizarPersonalidade).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", r))

}
