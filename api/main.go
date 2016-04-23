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

func Start(c config.Configuration) {
	log.Printf("Start api system on %s:%d", c.RConf.Admin.Bind, c.RConf.Admin.Port)
	r := mux.NewRouter()
	r.HandleFunc("/ping", PingHandler()).Methods("GET")
	r.HandleFunc("/backup", BackupHandler(c)).Methods("GET")
	r.HandleFunc("/frontend", GetFrontendsHandler(c)).Methods("GET")
	r.HandleFunc("/frontend/{name}", GetFrontendHandler(c)).Methods("GET")
	r.HandleFunc("/frontend/{name}/node", PostNodeHandler(c)).Methods("POST")
	r.HandleFunc("/frontend/{name}/node", DeleteNodeHandler(c)).Methods("DELETE")
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", c.RConf.Admin.Bind, c.RConf.Admin.Port), r)
	if err != nil {
		log.Fatalln(err)
	}
}
