package config

type RConf struct {
	Admin bool   `json:"admin"`
	Bind  string `json:"bind"`
	Port  int    `json:"port"`
}
