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

	DataBase DataBaseConfig `json:"database"`
}
type DataBaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	Charset  string `json:"charset"`
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
