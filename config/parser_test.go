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
