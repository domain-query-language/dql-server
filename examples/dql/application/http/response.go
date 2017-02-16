package http

import (
	"net/http"
	"encoding/json"
)

type responseFormat struct {
	data interface{}
}

func respondWithJson(w http.ResponseWriter, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response, _ := json.Marshal(responseFormat{data})
	w.Write(response)
}
