package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName string `json:"app_name"`

	AppMode string `json:"app_mode"`

	AppHost string `json:"app_host"`

	AppPort string `json:"app_port"`
}

func ParseConfig(path string) (*Config, error) {

	file, e := os.Open(path)

	if e != nil {
		panic(e)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	var config Config
	decode := decoder.Decode(&config)
	if decode != nil {
		return nil, decode
	}

	return &config, nil

}
