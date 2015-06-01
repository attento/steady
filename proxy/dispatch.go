package proxy

import (
	"math/rand"
	"net/http"
)

func CreateNewRequest(req *http.Request, nodes []string) *http.Request {
	newRequest := req
	server := nodes[rand.Intn(len(nodes))]
	newRequest.URL.Host = server
	newRequest.URL.Scheme = "http"
	return newRequest
}

func HydrateHeaders(resp *http.Response, w http.ResponseWriter) {
	for k, v := range resp.Header {
		for _, single := range v {
			w.Header().Set(k, single)
		}
	}
}
