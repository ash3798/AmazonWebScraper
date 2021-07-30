package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	ScraperPort       int    `split_words:"true" default:"9991"`
	PersistServiceURL string `split_words:"true" default:"http://localhost:9992"`
	mongodbHostname   string `split_words:"true" default:"localhost"`
	mongodbPort       int    `split_words:"true" default:"27071"`
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

// func GetMongoURL() string {
// 	uri := fmt.Sprintf("mongodb://%s:%d", Manager.mongodbHostname, Manager.mongodbPort)
// 	return uri
// }
