package handlers

import "net/http"

func HandleHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("test"))
}
