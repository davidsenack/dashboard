package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type metrics struct {
	Impressions int64 `json:"impressions"`
	Clicks      int64 `json:"clicks"`
	Conversions int64 `json:"conversions"`
}

// ::TODO Add Test Campaign Data
//type campaigns struct {
//	Campaigns
//}

// !!TODO:: Remove Fake Testing Data!
var fakeMetrics = []metrics{
	{Impressions: 23832, Clicks: 8244, Conversions: 1281},
	{Impressions: 374728, Clicks: 4728, Conversions: 918},
	{Impressions: 4785, Clicks: 1432, Conversions: 234},
	{Impressions: 4858211, Clicks: 38114, Conversions: 345},
	{Impressions: 79123, Clicks: 17274, Conversions: 542},
}

func getMetrics(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, fakeMetrics)
}

func main() {
	router := gin.Default()
	router.GET("/metrics", getMetrics)

	router.Run("localhost:8080")
}
