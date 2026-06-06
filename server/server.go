package server

import (
	"log"
	"net/http"
	api "notes/api"
)

func createServer() *http.ServeMux {
	var mux *http.ServeMux = http.NewServeMux()
	return mux
}

func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf(
			"%s %s %s",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
		)

		next.ServeHTTP(w, r)
	}
}

func StartAPIService() error {
	getRoot := logger(api.GetRoot)
	getHealth := logger(api.GetHealth)
	getNote := logger(api.GetNote)
	getAllNotes := logger(api.GetAllNotes)
	addNote := logger(api.AddNote)
	updateNote := logger(api.UpdateNote)
	deleteNote := logger(api.DeleteNote)
	generateApiKey := logger(api.GenerateAPIKey)

	server := createServer()
	server.HandleFunc("GET /", getRoot)
	server.HandleFunc("GET /healthz", getHealth)
	server.HandleFunc("GET /allnotes", getAllNotes)
	server.HandleFunc("GET /notes/{id}", getNote)
	server.HandleFunc("POST /notes/{id}/{content}", addNote)
	server.HandleFunc("PUT /notes/{id}/{content}", updateNote)
	server.HandleFunc("DELETE /notes/{id}", deleteNote)
	server.HandleFunc("GET /generate_api_key", generateApiKey)
	return http.ListenAndServe(":8080", server)
}
