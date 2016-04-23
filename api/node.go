package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gianarb/lb/config"
	"github.com/gianarb/lb/core"
)

func PostNodeHandler(config config.Configuration) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var server core.Server
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&server)
		if err != nil {
		}
		w.Header().Set("Content-Type", "application/json")
		h := strings.Split(r.URL.Path, "/")
		c := h[2]
		for key, val := range config.Frontends {
			if key == c {
				val.Nodes = append(val.Nodes, server)
				js, _ := json.Marshal(val)
				w.Write(js)
				return
			} else {
				w.WriteHeader(404)
			}
		}
	}
}

func DeleteNodeHandler(config config.Configuration) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var server core.Server
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&server)
		if err != nil {
			w.WriteHeader(406)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		h := strings.Split(r.URL.Path, "/")
		c := h[2]
		for key, val := range config.Frontends {
			if key == c {
				for k, nn := range val.Nodes {
					if nn.Host == server.Host {
						val.Nodes = append(val.Nodes[:k], val.Nodes[k+1:]...)
					}
				}
				js, _ := json.Marshal(val)
				w.Write(js)
				return
			} else {
				w.WriteHeader(404)
			}
		}
	}
}
