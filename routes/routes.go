package routes

import (
	"log"
	"net/http"

	"github.com/cledson-leite/go_api_rest.git/controllers"
)

func HandleRequest()  {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.ListarPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
