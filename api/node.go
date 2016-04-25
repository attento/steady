package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gianarb/steady/config"
	"github.com/gianarb/steady/core"
)

func PostNodeHandler(config config.Configuration) func(w http.ResponseWriter, r *http.Request) {
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
		fr := config.GetFrontendByName(c)
		if fr == nil {
			w.WriteHeader(404)
			return
		}
		fr.Nodes = append(fr.Nodes, server)
		js, _ := json.Marshal(fr)
		w.Write(js)
		return
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
		fr := config.GetFrontendByName(c)
		if fr == nil {
			w.WriteHeader(404)
			return
		}
		for k, nn := range fr.Nodes {
			if nn.Host == server.Host {
				fr.Nodes = append(fr.Nodes[:k], fr.Nodes[k+1:]...)
			}
		}
		js, _ := json.Marshal(fr)
		w.Write(js)
		return
	}
}
