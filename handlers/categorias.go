package handlers

import (
	"api/common"
	"api/db"
	"api/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ResponseGenerico struct {
	Estado  string
	Mensaje string
}

func Categoria_Get(response http.ResponseWriter, request *http.Request) {
	common.JSON_FORMAT(response)
	datos := model.Categorias{}
	db.Database.Order("id asc").Find(&datos)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(datos)

}

func Categoria_Get_Id(response http.ResponseWriter, request *http.Request) {
	common.JSON_FORMAT(response)
	vars := mux.Vars(request)
	// id , _ := strconv.Atoi(vars["id"])
	datos := model.Categoria{}

	if err := db.Database.First(&datos, vars["id"]); err.Error != nil {
		respuesta := map[string]string{
			"estado":  "ok",
			"mensaje": "Recurso no disponible",
		}

		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(respuesta)
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(datos)
	}

}
