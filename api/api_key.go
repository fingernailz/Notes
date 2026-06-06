package api

import "net/http"

func GenerateAPIKey(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("KYS"))
}
