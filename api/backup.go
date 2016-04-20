package api

import (
	"encoding/json"
	"net/http"

	"github.com/gianarb/lb/config"
)

func BackupHandler(conf config.Configuration) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		js, _ := json.Marshal(conf)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
