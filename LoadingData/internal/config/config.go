package config

import (
	"os"
)

type Config struct {
	ElasticsearchURL string
}

func LoadConfig() *Config {
	return &Config{
		ElasticsearchURL: os.Getenv("ELASTICSEARCH_URL"),
	}
}
