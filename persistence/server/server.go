package server

import (
	"github.com/ash3798/AmazonWebScraper/persistence/task"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

//Mapping function does the mapping of api's in router
func Mapping() {
	Router = gin.Default()

	api := Router.Group("/url")
	{
		api.POST("/persist", func(ctx *gin.Context) {
			productInfo := task.ProductInfo{}

			if err := ctx.ShouldBindJSON(&productInfo); err != nil {
				ctx.JSON(400, gin.H{
					"errorMessage": "Invalid payload. Cant read product Info. Error : " + err.Error(),
				})
			}

			if err := task.PersistDataToDB(productInfo); err != nil {
				ctx.JSON(500, gin.H{
					"errorMessage": "Not able to store the product info to database. Error : " + err.Error(),
				})
			}

			ctx.JSON(200, gin.H{
				"message": "product info stored successfully to database",
			})

		})
	}
}
