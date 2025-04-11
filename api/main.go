package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	address := "0.0.0.0:8080"
	r := getRouter(&DataFake{})
	r.Run(address)

}

func getRouter(db DataProvider) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1.GET("group/:id", func(c *gin.Context) {
		id := c.Param("id")
		home := db.getHouse(id)
		c.JSON(http.StatusOK, home)
	})
	v1.GET("water/:id", func(c *gin.Context) {
		id := c.Param("id")
		waterLog := db.getLatestWaterLogsForHome(id)
		c.JSON(http.StatusOK, waterLog)
	})
	v1.POST("plant/:id", func(c *gin.Context) {
		id := c.Param("id")
		db.wateredPlantAt(id, time.Now())
		c.Status(http.StatusOK)
	})
	return r
}
