package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gianarb/lb/core"
)

func (c *Configuration) Parse(filePath string) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &c)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

type Configuration struct {
	RConf     RConf                     `json:"rconf"`
	Frontends map[string]*core.Frontend `json:"frontends"`
}
