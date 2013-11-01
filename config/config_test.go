package config

import (
	"net"
	"testing"
)

func TestConfig(t *testing.T) {
	config, err := Parse()
	if err != nil {
		t.Errorf("Error occured: %v", err)
	}
	expected := net.ParseIP("127.0.0.1")

	if !expected.Equal(config.BoundIP()) {
		t.Errorf("Expected 127.0.0.1, got: %v", config.BoundIP())
	}
}
