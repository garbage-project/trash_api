package config

import (
	"github.com/naoina/toml"
	"os"
)

type DB struct {
	Database string
	Host     string
	User     string
	Password string

	Collections map[string]string
}

type Config struct {
	ServerInfo struct {
		Port    string
		Service string
		Log     string
	}

	DB map[string]DB
}

func NewConfig(file string) *Config {
	c := new(Config)

	if f, err := os.Open(file); err != nil {
		panic(err)
	} else {
		if err = toml.NewDecoder(f).Decode(c); err != nil {
			panic(err)
		} else {
			return c
		}
	}
}
