package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Igor0155/golang/go_api/database"
	"github.com/Igor0155/golang/go_api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Welcome to the Home Page!")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {

	var p []models.Personalidade

	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)
}

func PersonalidadePorId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade

	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)

}

func CriarPersonalidade(w http.ResponseWriter, r *http.Request) {
	var p models.Personalidade

	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Create(&p)

	json.NewEncoder(w).Encode(p)

}

func DeletarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade

	database.DB.Delete(&p, id)

	json.NewEncoder(w).Encode("Personalidade deletada com sucesso!")
}

func AtualizarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade

	database.DB.First(&p, id)

	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)

	json.NewEncoder(w).Encode(p)
}
