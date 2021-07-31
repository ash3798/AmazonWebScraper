package config

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	PersistServicePort  int    `split_words:"true" default:"9992"`
	MongodbHostname     string `split_words:"true" default:"localhost"`
	MongodbPort         int    `split_words:"true" default:"27071"`
	MongoDBName         string `split_words:"true" default:"product"`
	MongoCollectionName string `split_words:"true" default:"shopitems"`
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

func GetMongoURL() string {
	uri := fmt.Sprintf("mongodb://%s:%d", Manager.MongodbHostname, Manager.MongodbPort)
	log.Println("connecting to MongoAB URL :" + uri)
	return uri
}
