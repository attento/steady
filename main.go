package main

import (
	"fmt"
	"github.com/gianarb/lb/brain"
	"github.com/gianarb/lb/config"
	"github.com/gianarb/lb/proxy"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var conf config.Configuration
var db brain.Datastore

func run(w http.ResponseWriter, req *http.Request) {
	db.GetEnableInstance()
	newRequest := proxy.CreateNewRequest(req, "www.google.com")

	resp, err := http.Get(newRequest.URL.String())
	if err != nil {
		fmt.Printf("$s", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("$s", err)
	}

	proxy.HydrateHeaders(resp, w)
	w.Write([]byte(body))
}

func main() {
	conf.Parse("./lb.config.json")

	db := new(brain.Datastore)
	db.Open("./database")
	err := db.Init(conf.Nodes)

	if err != nil {
		fmt.Printf("%s \n", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/{[*]}", run)
	http.ListenAndServe(fmt.Sprintf(":%d", 8089), r)
}
