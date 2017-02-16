package http

import (
	"net/http"
	"encoding/json"
)

type success struct {
	Data interface{} `json:"data"`
}

type error struct {
	Error interface{} `json:"error"`
}

func respondWithJson(w http.ResponseWriter, data interface{}) {

	if (data == nil) {
		w.Header().Add("error", "No data")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(success{data})

	if (err != nil) {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, data interface{}) {

	response, err := json.Marshal(error{data})

	if (err != nil) {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(response)
}
