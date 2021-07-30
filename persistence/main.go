package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ash3798/AmazonWebScraper/persistence/config"
	"github.com/ash3798/AmazonWebScraper/persistence/server"
	"github.com/ash3798/AmazonWebScraper/persistence/task"
)

func main() {
	log.Println("Starting persistence Service")

	if !task.InitDatabaseClient() {
		return
	}

	if !config.InitEnv() {
		return
	}

	server.Mapping()

	addr := fmt.Sprintf(":%d", config.Manager.PersistServicePort)
	server.Router.Run(addr)

	err := task.MongoDB.Disconnect(context.TODO())
	if err != nil {
		log.Println("error while closing db connection , Error : ", err.Error())
	} else {
		log.Println("Connection to MongoDB closed.")
	}
}
