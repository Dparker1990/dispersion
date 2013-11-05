package config

import (
	"log"
	"net"
)

const configPath = "../config.json"

type Config struct {
	Bindip   string
	boundip  net.IP
	Bindport string
	Seeds    []string
	Logfile  string
}

func (c Config) BoundIP() (ip net.IP) {
	if ip = c.boundip; ip != nil {
		return
	}

	if ip = net.ParseIP(c.Bindip); ip == nil {
		log.Fatalf("Could not parse IP, got: %v, original was: %v", ip, c.Bindip)
	}

	c.boundip = ip

	return
}
