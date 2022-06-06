package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(write http.ResponseWriter, response *http.Request)  {
	fmt.Fprint(write, "Home page")
}

func HandleRequest()  {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	fmt.Println("Olá Mundo!!")
	HandleRequest()
}