package common

import (
	"net/http"
)

func JSON_FORMAT(response http.ResponseWriter) {

	response.Header().Set("Content-type", "application/json")

}
