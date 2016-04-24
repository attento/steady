package config

import (
	"testing"
)

func TestParseWithCorrectPath(t *testing.T) {
	_, err := Parse("../lb.config.json")
	if err != nil {
		t.Fail()
	}
}

func TestParseWrongPath(t *testing.T) {
	_, err := Parse("../random.path")
	if err == nil {
		t.Fail()
	}
}

func TestGetFrontendByName(t *testing.T) {
	c, _ := Parse("../lb.config.json")
	dd := c.GetFrontendByName("andrea")
	if dd == nil {
		t.Fail()
	}
}

func TestGetUnexistedFrontendByName(t *testing.T) {
	c, _ := Parse("../lb.config.json")
	dd := c.GetFrontendByName("notexist")
	if dd != nil {
		t.Fail()
	}
}
