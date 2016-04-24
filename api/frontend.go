package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gianarb/lb/config"
	"github.com/gianarb/lb/core"
	"github.com/gianarb/lb/proxy"
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
		dd := make(chan bool)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
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
		go func() {
			log.Printf("Start %s on %s:%d", c, frontend.Bind, frontend.Port)
			err := http.ListenAndServe(fmt.Sprintf("%s:%d", frontend.Bind, frontend.Port), proxy.ProxyHandler(&frontend))
			if err != nil {
				dd <- false
				return
			}
			dd <- true
		}()
		for i := range dd {
			log.Println("Start")
			if i == true {
				w.WriteHeader(200)
				js, _ := json.Marshal(config.Frontends[c])
				w.Write(js)
				return
			} else {
				w.WriteHeader(500)
				return
			}
		}
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
