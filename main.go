package main

import (
	"flag"
	"log"
	"os"
	"sync"

	"github.com/gianarb/lb/api"
	"github.com/gianarb/lb/config"
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
	for name, v := range c.Frontends {
		wg.Add(1)

		go proxy.StartFrontend(name, v)
	}
	wg.Wait()
}
