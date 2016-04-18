package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gianarb/lb/config"
	"github.com/gorilla/mux"
)

type Api struct {
}

func Start(c config.RConf) {
	log.Printf("Start api system on %s:%d", c.Bind, c.Port)
	r := mux.NewRouter()
	r.HandleFunc("/ping", PingHandler()).Methods("GET")
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", c.Bind, c.Port), r)
	if err != nil {
		log.Fatalln(err)
	}
}
