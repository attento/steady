package core

type Server struct {
	Host   string            `json:"host"`
	Fields map[string]string `json:"fields"`
}
