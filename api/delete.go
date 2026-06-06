package api

import (
	"net/http"
	"strconv"
)

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	id, z := strconv.Atoi(r.PathValue("id"))

	if z != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if !SetOfBooks.DeleteBook(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusAccepted)
}
