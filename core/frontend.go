package core

type Frontend struct {
	Port  int      `json:"port"`
	Bind  string   `json:"bind"`
	Nodes []Server `json:"nodes"`
}
