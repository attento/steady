package config

import "github.com/gianarb/steady/core"

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
