package api

import (
	"encoding/json"
	"net/http"
	"strings"

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
		h := strings.Split(r.URL.Path, "/")
		c := h[2]
		for key, val := range config.Frontends {
			if key == c {
				js, _ := json.Marshal(val)
				w.WriteHeader(200)
				w.Write(js)
				return
			} else {
				w.WriteHeader(404)
			}
		}
	}
}
