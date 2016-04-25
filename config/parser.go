package config

import (
	"encoding/json"
	"io/ioutil"
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
