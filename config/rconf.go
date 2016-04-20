package config

type Bindable struct {
	Bind string `json:"bind"`
	Port int    `json:"port"`
}

type RConf struct {
	Admin *Bindable `json:"admin,omitempty"`
	Ui    *Bindable `json:"ui,omitempty"`
}
