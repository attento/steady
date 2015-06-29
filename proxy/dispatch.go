package proxy

import (
	"net/http"
)

func CreateNewRequest(req *http.Request, node string) *http.Request {
	newRequest := req
	newRequest.URL.Host = node
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
