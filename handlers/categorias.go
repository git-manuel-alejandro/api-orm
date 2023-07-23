package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseGenerico struct {
	Estado  string
	Mensaje string
}

func Categoria_Get(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-type", "application/json")
	output, _ := json.Marshal(ResponseGenerico{"ok", "GET"})
	fmt.Fprintln(response, string(output))

}
