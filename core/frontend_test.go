package core

import "testing"

func TestDeleteNode(t *testing.T) {
	var nodes []Server
	nodes = append(
		nodes,
		Server{Host: "127.0.0.1"},
		Server{Host: "127.0.0.10"},
		Server{Host: "127.0.0.50"},
	)
	fr := Frontend{
		Port:  8080,
		Bind:  "0.0.0.0:8080",
		Nodes: nodes,
	}
	fr.DeleteNodeByHost("127.0.0.10")
	for _, v := range fr.Nodes {
		if v.Host == "127.0.0.10" {
			t.Fail()
		}
	}
}

func TestDeleteLastNode(t *testing.T) {
	var nodes []Server
	nodes = append(
		nodes,
		Server{Host: "127.0.0.1"},
		Server{Host: "127.0.0.10"},
	)
	fr := Frontend{
		Port:  8080,
		Bind:  "0.0.0.0:8080",
		Nodes: nodes,
	}
	fr.DeleteNodeByHost("127.0.0.10")
	for _, v := range fr.Nodes {
		if v.Host == "127.0.0.10" {
			t.Fail()
		}
	}
}

func TestAddNode(t *testing.T) {
	var nodes []Server
	nodes = append(
		nodes,
		Server{Host: "127.0.0.12"},
		Server{Host: "127.0.0.10"},
	)
	fr := Frontend{
		Port:  8080,
		Bind:  "0.0.0.0:8080",
		Nodes: nodes,
	}
	fr.AddNode(Server{Host: "127.0.0.1"})
	check := false
	for _, v := range fr.Nodes {
		if v.Host == "127.0.0.1" {
			check = true
		}
	}
	if check != true {
		t.Fail()
	}
}
