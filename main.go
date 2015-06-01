package main

import (
	"fmt"
	"github.com/gianarb/lb/config"
	"github.com/gianarb/lb/proxy"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var conf config.Configuration

func run(w http.ResponseWriter, req *http.Request) {
	newRequest := proxy.CreateNewRequest(req, conf.Nodes)

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
	r := mux.NewRouter()
	r.HandleFunc("/{[*]}", run)
	http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), r)
}
