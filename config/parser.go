package config

import (
	"encoding/json"
	"io/ioutil"
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

// Configuration configuration struct
type Configuration struct {
	Port  int      `json:"port"`
	Nodes []string `json:"nodes"`
}
