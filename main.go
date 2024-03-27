package main

import (
	"net/http"

	"github.com/erranteurbano/api-go-resfull/routes"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)
}
