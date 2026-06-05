package api

import (
	"net/http"
)

func GetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func GetRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HI THIS IS THE NOTES API"))
}

func GetAllNotes(w http.ResponseWriter, r *http.Request) {

}
