package api

import (
	"encoding/json"
	"net/http"

	"github.com/gianarb/steady/config"
)

func BackupHandler(conf config.Configuration) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		js, err := json.Marshal(conf)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
		w.Write(js)
	}
}
