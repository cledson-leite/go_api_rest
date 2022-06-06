package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cledson-leite/go_api_rest.git/models"
	"github.com/gorilla/mux"
)

func Home(write http.ResponseWriter, request *http.Request)  {
	fmt.Fprint(write, "Home page")
}

func ListarPersonalidades(write http.ResponseWriter, request *http.Request){
	json.NewEncoder(write).Encode(models.Personalidades)
}

func BuscarPorId(write http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	for _, personalidade := range models.Personalidades {
		if strconv.Itoa(personalidade.Id) == id {
			json.NewEncoder(write).Encode(personalidade)
		}
	}
}