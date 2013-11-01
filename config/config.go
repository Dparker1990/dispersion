package config

import (
	"encoding/json"
	"io/ioutil"
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

func Parse() (conf Config, err error) {
	contents, err := loadConfig()
	if err != nil {
		log.Fatalf("Could not load config due to: %v", err)
	}

	conf = parseConfig(contents)

	return
}

func loadConfig() (contents []byte, err error) {
	contents, err = ioutil.ReadFile(configPath)
	return
}

func parseConfig(contents []byte) (c Config) {
	err := json.Unmarshal(contents, &c)
	if err != nil {
		log.Fatalf("Could not parse config due to: %v", err)
	}
	return
}
