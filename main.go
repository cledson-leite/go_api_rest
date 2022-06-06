package main

import (
	"fmt"

	"github.com/cledson-leite/go_api_rest.git/models"
	"github.com/cledson-leite/go_api_rest.git/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Olá Mundo!!")
	routes.HandleRequest()
}