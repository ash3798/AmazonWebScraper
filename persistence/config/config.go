package config

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	PersistServicePort  int    `split_words:"true" default:"9992"`
	MongodbHostname     string `split_words:"true" default:"localhost"`
	MongodbPort         int    `split_words:"true" default:"27017"`
	MongoDBName         string `split_words:"true" default:"product"`
	MongoCollectionName string `split_words:"true" default:"shopitems"`
}

var (
	Manager config
)

//InitEnv function reads the environment variables and initialize the config variable with them
func InitEnv() bool {
	err := envconfig.Process("", &Manager)
	if err != nil {
		log.Println("Error while initializing environment")
		return false
	}
	return true
}

//GetMongoURL function creates the mongo db URI according to the mentioned hostname and port
func GetMongoURL() string {
	uri := fmt.Sprintf("mongodb://%s:%d", Manager.MongodbHostname, Manager.MongodbPort)
	log.Println("connecting to MongoDB URL :" + uri)
	return uri
}
