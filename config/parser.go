package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gianarb/lb/redundancy"
)

func (c *Configuration) Parse(filePath string) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &c)
	if err != nil {
		return err
	}
	return nil
}

type Configuration map[string]*redundancy.Frontend
