package helpers

import (
	"encoding/json"
	"net/http"
)

func Ok(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(body)
	return
}

func BadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(err)
	return
}

func NotFound(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(err)
	return
}
