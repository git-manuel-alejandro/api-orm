package main

import (
	"api/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// model.Migraciones()
	mux := mux.NewRouter()
	prefijo := "/api/v1/"

	mux.HandleFunc(prefijo+"categorias", handlers.Categoria_Get).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", mux))

}
