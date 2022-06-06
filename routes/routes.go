package routes

import (
	"log"
	"net/http"

	"github.com/cledson-leite/go_api_rest.git/controllers"
)

func HandleRequest()  {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
