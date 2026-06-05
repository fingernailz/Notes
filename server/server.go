package server

import (
	"net/http"
	api "notes/api"
)

func createServer() *http.ServeMux {
	var mux *http.ServeMux = http.NewServeMux()
	return mux
}

func StartAPIService() error {
	server := createServer()
	server.HandleFunc("GET /", api.GetRoot)
	server.HandleFunc("GET /healthz", api.GetHealth)
	server.Handle("GET /notes", api.GetAllNotes)
	return http.ListenAndServe(":8080", server)
}
