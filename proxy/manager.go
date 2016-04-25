package proxy

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gianarb/steady/core"
)

func StartFrontend(name string, frontend *core.Frontend) error {
	server := &http.Server{Handler: ProxyHandler(frontend)}
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", frontend.Bind, frontend.Port))
	defer listener.Close()
	if nil != err {
		return err
	}
	log.Printf("Start %s on %s:%d", name, frontend.Bind, frontend.Port)
	if err := server.Serve(listener); nil != err {
		return err
	}
	return nil
}
