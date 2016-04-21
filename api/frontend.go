package api

import (
	"encoding/json"
	"net/http"

	"github.com/gianarb/lb/config"
)

type FrontendResponse struct {
}

func GetFrontendsHandler(config config.Configuration) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		js, _ := json.Marshal(config.Frontends)
		w.WriteHeader(200)
		w.Write(js)
	}
}

func GetFrontendHandler(config config.Configuration) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		c := "andrea"
		for key, val := range config.Frontends {
			if key == c {
				js, _ := json.Marshal(val)
				w.WriteHeader(200)
				w.Write(js)
			} else {
				w.WriteHeader(404)
			}
		}
	}
}
