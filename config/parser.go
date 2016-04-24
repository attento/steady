package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gianarb/lb/core"
)

func Parse(filePath string) (Configuration, error) {
	var c Configuration
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(file, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}

type Configuration struct {
	RConf     RConf                     `json:"rconf"`
	Frontends map[string]*core.Frontend `json:"frontends"`
}

func (c *Configuration) GetFrontendByName(name string) *core.Frontend {
	for key, val := range c.Frontends {
		if key == name {
			return val
		}
	}
	return nil
}
