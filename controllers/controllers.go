package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/cledson-leite/go_api_rest.git/database"
	"github.com/cledson-leite/go_api_rest.git/models"
	"github.com/gorilla/mux"
)

func Home(write http.ResponseWriter, request *http.Request)  {
	fmt.Fprint(write, "Home page")
}

func ListarPersonalidades(write http.ResponseWriter, request *http.Request){
	var personalidades []models.Personalidade
	database.DB.Find(&personalidades)

	json.NewEncoder(write).Encode(personalidades)
}

func BuscarPorId(write http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(write).Encode(personalidade)

}