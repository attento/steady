package core

type Frontend struct {
	Port  int      `json:"port"`
	Bind  string   `json:"bind"`
	Nodes []Server `json:"nodes"`
}

func (fr *Frontend) DeleteNodeByHost(host string) {
	for k, nn := range fr.Nodes {
		if nn.Host == host {
			fr.Nodes = append(fr.Nodes[:k], fr.Nodes[k+1:]...)
		}
	}
}

func (fr *Frontend) AddNode(server Server) {
	fr.Nodes = append(fr.Nodes, server)
}
