package proxy

import (
	"math/rand"
	"net/http"

	"github.com/gianarb/lb/redundancy"
)

func CreateNewRequest(req *http.Request, nodes []redundancy.Server) *http.Request {
	newRequest := req
	server := nodes[rand.Intn(len(nodes))]
	newRequest.URL.Host = server.Host
	newRequest.URL.Scheme = "http"
	return newRequest
}

func HydrateHeaders(resp *http.Response, w http.ResponseWriter) {
	w.WriteHeader(resp.StatusCode)
	for k, v := range resp.Header {
		for _, single := range v {
			w.Header().Set(k, single)
		}
	}
}
