package api

import (
	"strconv"
	"suraj_projects/allen_interview/deal"

	"github.com/gin-gonic/gin"
)

func CreateNewUserController(ctx *gin.Context) {
	var user deal.User
	ctx.BindJSON(&user)
	u := *user.NewUser(user.UserName, user.Email)

	ctx.JSON(200, gin.H{
		"message": "User Created with id: " + strconv.Itoa(u.Id),
	})
	// Create a new user
}

type ClaimDealStruct struct {
	UserId     int `json:"user_id"`
	DealId     int `json:"deal_id"`
	Product_id int `json:"product_id"`
}

func ClaimDealController(ctx *gin.Context) {
	var u deal.User
	var p deal.Product
	ClaimDealStruct := ClaimDealStruct{}
	ctx.BindJSON(&ClaimDealStruct)

	// Claim a deal
	product := p.GetProduct(ClaimDealStruct.Product_id)
	user := u.GetUser(ClaimDealStruct.UserId)
	_, err := user.ClaimDeal(product)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "Deal Claimed",
		})

	}

}

type UserId struct {
	UserId int `json:"user_id"`
}

func GetUserController(ctx *gin.Context) {
	var gu UserId

	ctx.BindJSON(&gu)
	var user *deal.User
	user = user.GetUser(gu.UserId)
	ctx.JSON(200, gin.H{
		"message": user,
	})
}
