package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	ScraperServicePort int    `split_words:"true" default:"9991"`
	PersistServiceURL  string `split_words:"true" default:"http://localhost:9992/url/persist"`
}

var (
	Manager config
)

func InitEnv() bool {
	err := envconfig.Process("", &Manager)
	if err != nil {
		log.Println("Error while initializing environment")
		return false
	}
	return true
}
