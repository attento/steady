package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gianarb/lb/api"
	"github.com/gianarb/lb/config"
	"github.com/gianarb/lb/core"
	"github.com/gianarb/lb/proxy"
)

func main() {
	var configPath string
	cmdFlags := flag.NewFlagSet("event", flag.ContinueOnError)
	cmdFlags.StringVar(&configPath, "c", "/etc/lb.config.json", "c")

	if err := cmdFlags.Parse(os.Args[1:]); err != nil {
		log.Fatalln(err)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalln(err)
	}

	c, err := config.Parse(configPath)
	if err != nil {
		log.Fatalln(err)
	}
	if c.RConf.Admin != nil {
		go api.Start(c)
	}

	var wg sync.WaitGroup
	for name, frontend := range c.Frontends {
		wg.Add(1)
		go func(fr *core.Frontend, n string) {
			defer wg.Done()
			log.Printf("Start %s on %s:%d", n, fr.Bind, fr.Port)
			err := http.ListenAndServe(fmt.Sprintf("%s:%d", fr.Bind, fr.Port), proxy.ProxyHandler(fr))
			if err != nil {
				log.Fatalln(err)
			}
		}(frontend, name)
	}
	wg.Wait()
}
