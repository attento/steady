package proxy

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/gianarb/steady/core"
)

func createNewRequest(req *http.Request, nodes []core.Server) *http.Request {
	newRequest := req
	server := nodes[rand.Intn(len(nodes))]
	newRequest.URL.Host = server.Host
	newRequest.URL.Scheme = "http"
	return newRequest
}

func hydrateHeaders(resp *http.Response, w http.ResponseWriter) {
	w.WriteHeader(resp.StatusCode)
	for k, v := range resp.Header {
		for _, single := range v {
			w.Header().Set(k, single)
		}
	}
}

func ProxyHandler(fr *core.Frontend) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		newRequest := createNewRequest(req, fr.Nodes)

		resp, err := http.Get(newRequest.URL.String())
		if err != nil {
			fmt.Printf("$s", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("$s", err)
		}
		hydrateHeaders(resp, w)
		w.Write([]byte(body))
	}
}
