package main

import (
	"fmt"
	"github.com/gianarb/lb/config"
	"github.com/gorilla/mux"
	"io/ioutil"
	"math/rand"
	"net/http"
)

var conf config.Configuration

func createNewRequest(req *http.Request) *http.Request {
	newRequest := req
	server := conf.Nodes[rand.Intn(len(conf.Nodes))]
	fmt.Printf("%s\n", server)
	newRequest.URL.Host = server
	newRequest.URL.Scheme = "http"
	return newRequest
}

func run(w http.ResponseWriter, req *http.Request) {
	newRequest := createNewRequest(req)

	resp, err := http.Get(newRequest.URL.String())
	if err != nil {
		fmt.Printf("$s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("$s", err)
	}
	for k, v := range resp.Header {
		for _, single := range v {
			w.Header().Set(k, single)
		}
	}
	w.Write([]byte(body))
}

func main() {
	conf.Parse("./lb.config.json")
	r := mux.NewRouter()
	r.HandleFunc("/{[*]}", run)
	http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), r)
}
