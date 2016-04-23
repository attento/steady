package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gianarb/lb/config"
	"github.com/gianarb/lb/core"
)

type FrontendResponse struct {
}

func DeleteFrontendsHandler(config config.Configuration) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		h := strings.Split(r.URL.Path, "/")
		c := h[2]
		delete(config.Frontends, c)
		w.WriteHeader(201)
	}
}

func PostFrontendsHandler(config config.Configuration) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var frontend core.Frontend
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&frontend)
		if err != nil {
			w.WriteHeader(406)
			return
		}
		h := strings.Split(r.URL.Path, "/")
		c := h[2]
		config.Frontends[c] = &frontend
		w.WriteHeader(200)
		js, _ := json.Marshal(config.Frontends[c])
		w.Write(js)
	}
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
