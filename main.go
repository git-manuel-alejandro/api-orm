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
	mux.HandleFunc(prefijo+"categorias/{id:[0-9]+}", handlers.Categoria_Get_Id).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", mux))

}
