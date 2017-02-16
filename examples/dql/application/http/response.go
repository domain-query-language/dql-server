package http

import (
	"net/http"
	"encoding/json"
)

type responseFormat struct {
	Data interface{} `json:"data"`
}

func respondWithJson(w http.ResponseWriter, data interface{}) {

	if (data == nil) {
		w.Header().Add("error", "No data")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(responseFormat{data})

	if (err != nil) {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
