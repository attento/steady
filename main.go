package main

import (
	"fmt"
	"github.com/gianarb/lb/config"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func run(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s\n", req.URL.Scheme)
	newRequest := req
	newRequest.URL.Host = "localhost:8080"
	newRequest.URL.Scheme = "http"
	resp, err := http.Get(newRequest.URL.String())
	if err != nil {
		fmt.Printf("$s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("$s", err)
	}
	w.Write([]byte(body))
}

func main() {
	var conf config.Configuration
	conf.Parse("./lb.config.json")
	r := mux.NewRouter()
	r.HandleFunc("/{[*]}", run)
	http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), r)
}
