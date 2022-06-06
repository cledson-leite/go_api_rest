package controllers

import (
	"fmt"
	"net/http"
)

func Home(write http.ResponseWriter, response *http.Request)  {
	fmt.Fprint(write, "Home page")
}