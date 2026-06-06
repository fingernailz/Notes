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
	server.HandleFunc("GET /allnotes", api.GetAllNotes)
	server.HandleFunc("GET /notes/{id}", api.GetNote)
	server.HandleFunc("POST /notes/{id}/{content}", api.AddNote)
	server.HandleFunc("PUT /notes/{id}/{content}", api.UpdateNote)
	server.HandleFunc("DELETE /notes/{id}", api.DeleteNote)
	server.HandleFunc("GET /generate_api_key", api.GenerateAPIKey)
	return http.ListenAndServe(":8080", server)
}
