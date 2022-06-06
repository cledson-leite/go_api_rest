package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cledson-leite/go_api_rest.git/models"
)

func Home(write http.ResponseWriter, response *http.Request)  {
	fmt.Fprint(write, "Home page")
}

func ListarPersonalidades(write http.ResponseWriter, response *http.Request){
	json.NewEncoder(write).Encode(models.Personalidades)
}