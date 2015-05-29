package main

import (
	"fmt"
	"github.com/gianarb/lb/config"
	"github.com/gorilla/mux"
	"net/http"
)

func run(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("X-Lb", "Load-1")
}

func main() {
	var conf config.Configuration
	conf.Parse("./lb.config.json")
	r := mux.NewRouter()
	r.HandleFunc("/{[*]}", run)
	http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), r)
}
