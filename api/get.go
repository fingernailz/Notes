package api

import (
	"encoding/json"
	"net/http"
	storage "notes/storage"
	"strconv"
)

var SetOfBooks storage.Library

func init() {
	SetOfBooks = storage.Library{}
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func GetRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HI THIS IS THE NOTES API"))
}

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	if len(SetOfBooks) == 0 {
		w.WriteHeader(http.StatusNoContent)
	}

	notesBin, err := json.Marshal(SetOfBooks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(notesBin)
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	x, y := strconv.Atoi(r.PathValue("id"))
	if y != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	bookPtr := SetOfBooks.SearchBook(x)
	if bookPtr == nil {
		w.WriteHeader(http.StatusConflict)
	}

	var noteInterface map[string]interface{} = map[string]interface{}{
		"id":      x,
		"content": bookPtr.Content,
	}

	// noteJson, err := json.Marshal(noteInterface)

	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(noteInterface)
}
