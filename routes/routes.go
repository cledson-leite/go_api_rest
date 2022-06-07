package routes

import (
	"log"
	"net/http"

	"github.com/cledson-leite/go_api_rest.git/controllers"
	"github.com/cledson-leite/go_api_rest.git/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest()  {
	router := mux.NewRouter()
	router.Use(middleware.ContentTypeMiddleware)

	router.HandleFunc("/", controllers.Home)

	router.HandleFunc(
		"/api/personalidades", controllers.ListarPersonalidades,
		).Methods("Get")

	router.HandleFunc(
		"/api/personalidades/{id}", controllers.BuscarPorId,
		).Methods("Get")

	router.HandleFunc(
		"/api/personalidades", controllers.CriarPersonalidade,
		).Methods("Post")

	router.HandleFunc(
		"/api/personalidades/{id}", controllers.EditarPersonalidade,
		).Methods("Put")

	router.HandleFunc(
		"/api/personalidades/{id}", controllers.DeletarPersonalidade,
		).Methods("Delete")
		
	log.Fatal(
		http.ListenAndServe(
			":8000", handlers.CORS(
				handlers.AllowedOrigins(
					[] string {"*"},
				),
			)(router),
		),
	)
}
