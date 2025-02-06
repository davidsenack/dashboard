package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type metrics struct {
	Campaign    string `json:"campaign"`
	Impressions int64  `json:"impressions"`
	Clicks      int64  `json:"clicks"`
	Conversions int64  `json:"conversions"`
}

// ::TODO Add Test Campaign Data
//type campaigns struct {
//	Campaigns
//}

// !!TODO:: Remove Fake Testing Data!
var fakeMetrics = []metrics{
	{Campaign: "Save the Bees", Impressions: 23832, Clicks: 8244, Conversions: 1281},
	{Campaign: "Save the Trees", Impressions: 374728, Clicks: 4728, Conversions: 918},
	{Campaign: "Save the Clouds", Impressions: 4785, Clicks: 1432, Conversions: 234},
	{Campaign: "Destroy the Clouds", Impressions: 4858211, Clicks: 38114, Conversions: 345},
	{Campaign: "Fish", Impressions: 79123, Clicks: 17274, Conversions: 542},
}

func getMetrics(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, fakeMetrics)
}

func main() {
	router := gin.Default()
	router.GET("/metrics", getMetrics)

	router.Run("localhost:8080")
}
