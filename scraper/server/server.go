package server

import (
	"github.com/ash3798/AmazonWebScraper/scraper/task"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitMappings() {
	Router = gin.Default()

	api := Router.Group("/url")
	{
		api.POST("/scrape", func(ctx *gin.Context) {
			var urlInfo task.UrlInfo

			if err := ctx.ShouldBindJSON(&urlInfo); err != nil {
				ctx.JSON(400, gin.H{"error": "Could not read URL , " + err.Error()})
				return
			}

			res, err := task.ScrapeAndSend(urlInfo.Url)
			if err != nil {
				ctx.JSON(500, gin.H{
					"error": "Scrape not successful. Error : " + err.Error(),
				})
				return
			}

			ctx.Header("Content-Type", "application/json; charset=utf-8")
			ctx.String(200, string(res))
		})
	}
}
