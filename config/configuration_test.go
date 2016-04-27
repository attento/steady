package config

import (
	"testing"

	"github.com/gianarb/steady/core"
)

func TestGetFrontendByName(t *testing.T) {
	frs := make(map[string]*core.Frontend)
	frs["app1"] = &core.Frontend{
		Bind: "0.0.0.0",
		Port: 123,
	}
	c := Configuration{
		Frontends: frs,
	}
	f := c.GetFrontendByName("app1")
	if f == nil {
		t.Fail()
	}
}

func TestGetUnexistedFrontendByName(t *testing.T) {
	frs := make(map[string]*core.Frontend)
	frs["app1"] = &core.Frontend{
		Bind: "0.0.0.0",
		Port: 123,
	}
	c := Configuration{
		Frontends: frs,
	}
	f := c.GetFrontendByName("app2")
	if f != nil {
		t.Fail()
	}
}
