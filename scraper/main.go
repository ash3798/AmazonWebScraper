package main

import (
	"fmt"
	"log"

	"github.com/ash3798/AmazonWebScraper/scraper/config"
	"github.com/ash3798/AmazonWebScraper/scraper/server"
)

func main() {
	log.Println("Starting scrapper Service")

	if !config.InitEnv() {
		return
	}

	server.InitMappings()
	addr := fmt.Sprintf(":%d", config.Manager.ScraperPort)
	server.Router.Run(addr)
}
