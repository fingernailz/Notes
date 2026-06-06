package api

import (
	"net/http"
	types "notes/storage"
	"strconv"
	"strings"
)

func AddNote(w http.ResponseWriter, r *http.Request) {
	id, z := strconv.Atoi(r.PathValue("id"))
	content := r.PathValue("content")

	if z != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if SetOfBooks.SearchBook(id) != nil {
		w.WriteHeader(http.StatusConflict)
	}

	content = strings.ReplaceAll(content, "_", " ")
	var book types.Book = types.Book{
		Content: content,
	}
	SetOfBooks[types.Id(id)] = book
	w.WriteHeader(http.StatusCreated)
}
