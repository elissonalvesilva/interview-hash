package helpers

import (
	"encoding/json"
	presenterProtocols "github.com/elissonalvesilva/interview-hash/my-ecommerce/presenters/protocols"
	"net/http"
)

func Ok(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(body)
}

func BadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(presenterProtocols.ErrorResponse{
		Stack: err,
		Message: err.Error(),
	})
}

func NotFound(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(presenterProtocols.ErrorResponse{
		Stack: err,
		Message: err.Error(),
	})
}
