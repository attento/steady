package main

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func run(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("X-Lb", "Load-1")
	io.WriteString(w, "ciao")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{*}", run)
	http.ListenAndServe(":8080", r)
}
