package api

import (
	"strconv"
	"suraj_projects/allen_interview/deal"

	dealProperties "suraj_projects/allen_interview/deal"

	"github.com/gin-gonic/gin"
)

//

func CreateDealController(c *gin.Context) {
	var deal *deal.Deal
	c.BindJSON(&deal)
	deal = deal.NewDeal(deal.DealStartTime, deal.DealEndTime, deal.DealName, deal.DealDescription, deal.DealPrice, deal.MaxDealQuantity, deal.DealQuantity, deal.Products)
	dealProperties.DealsList = append(dealProperties.DealsList, deal)

	c.JSON(200, gin.H{
		"message": "Deal Created with id: " + strconv.Itoa(deal.Id),
	})
}

func UpdateDealController(c *gin.Context) {
	var deal *deal.Deal
	c.BindJSON(&deal)
	deal.UpdateDeal(deal.DealEndTime, deal.MaxDealQuantity)

	c.JSON(200, gin.H{
		"message": "Deal Updated",
	})
}

type GetDeal struct {
	Id int `json:"id"`
}

func GetDealController(c *gin.Context) {
	var gd GetDeal
	var deal *deal.Deal
	c.BindJSON(&gd)
	deal = deal.GetDeal(gd.Id)
	// return deal object in json response to client
	c.JSON(200, gin.H{
		"message": deal,
	})

}
