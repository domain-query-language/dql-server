package http

import (
	"net/http"
	"log"
)

const PORT = ":4242"

func StartHttpServer() {

	server := SetupServer()

	err := http.ListenAndServe(PORT, server)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func SetupServer() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/schema", Schema)

	return mux
}
