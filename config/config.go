package config

import (
	"encoding/json"
	"io/ioutil"
)

const configPath = "../config.json"

type Config struct {
	bindip      net.IP
	bindport    string
	constactip  net.IP
	contactport string
	protocol    string
	logfile     string
}

func Configuration() (c Config, err error) {
	contents, err := loadConfig()
	if err != nil {
		return
	}

	c, err = parseConfig(contents)

	return
}

func loadConfig() (contents []byte, err error) {
	contents, err = ioutil.ReadFile(configPath)
	return
}

func parseConfig(contents []byte) (c Config, err error) {
	err = json.Unmarshal(contents, &c)
	return
}
