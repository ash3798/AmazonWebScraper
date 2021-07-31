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
	//all config variables can be accessed through Manager
	Manager config
)

//InitEnv function reads the environments variables and initializes the config variables with them
func InitEnv() bool {
	err := envconfig.Process("", &Manager)
	if err != nil {
		log.Println("Error while initializing environment")
		return false
	}
	return true
}
