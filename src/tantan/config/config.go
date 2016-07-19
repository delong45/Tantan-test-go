package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	DefaultHTTPAddr string = ":8081"
)

type Config struct {
	HTTPAddr string
}

func Parse(configFile string) (*Config, error) {
	c, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("config.Parse read file error: %v", err)
	}

	var conf Config
	if err := json.Unmarshal(c, &conf); err != nil {
		return nil, fmt.Errorf("config.Parse json error: %v", err)
	}

	if conf.HTTPAddr == "" {
		conf.HTTPAddr = DefaultHTTPAddr
	}

	return &conf, nil
}
