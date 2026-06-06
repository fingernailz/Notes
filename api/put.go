package api

import (
	"net/http"
	"strconv"
)

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	id, z := strconv.Atoi(r.PathValue("id"))
	content := r.PathValue("content")

	if z != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	bookPtr := SetOfBooks.SearchBook(id)
	if bookPtr == nil {
		w.WriteHeader(http.StatusConflict)
	}

	bookPtr.Content = content
	w.WriteHeader(http.StatusAccepted)
}
