package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gianarb/steady/config"
	"github.com/gianarb/steady/core"
	"github.com/gianarb/steady/proxy"
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
		var frontend core.Frontend

		h := strings.Split(r.URL.Path, "/")
		c := h[2]

		w.Header().Set("Content-Type", "application/json")
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&frontend)
		config.Frontends[c] = &frontend
		if err != nil {
			w.WriteHeader(406)
			return
		}

		go proxy.StartFrontend(c, &frontend)
		w.WriteHeader(200)
		js, _ := json.Marshal(frontend)
		w.Write(js)
		return
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

func GetFrontendHandler(myConf config.Configuration) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h := strings.Split(r.URL.Path, "/")
		c := h[2]
		fr := myConf.GetFrontendByName(c)
		if fr == nil {
			w.WriteHeader(404)
			return
		}
		js, _ := json.Marshal(fr)
		w.WriteHeader(200)
		w.Write(js)
	}
}
