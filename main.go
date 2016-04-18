package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gianarb/lb/config"
	"github.com/gianarb/lb/proxy"
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
	http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), http.HandlerFunc(run))
}
