package api

import (
	"strconv"
	"suraj_projects/allen_interview/deal"

	"github.com/gin-gonic/gin"
)

func CreateNewProductController(c *gin.Context) {
	var product *deal.Product
	c.BindJSON(&product)
	product = product.NewProduct(product.ProductName, product.ProductDescription)
	deal.ProductsList = append(deal.ProductsList, product)

	c.JSON(200, gin.H{
		"message": "Product Created with id: " + strconv.Itoa(product.Id),
	})
}

type AssignDealToProduct struct {
	ProductId int `json:"product_id"`
	DealId    int `json:"deal_id"`
}

func AssignDealToProductController(c *gin.Context) {
	var p *deal.Product
	var d *deal.Deal
	// get id field from request json

	var assignDealToProduct AssignDealToProduct
	c.BindJSON(&assignDealToProduct)
	deal := d.GetDeal(assignDealToProduct.DealId)
	product := p.GetProduct(assignDealToProduct.ProductId)
	// assign deal to product
	product.AssignDealToProduct(deal)

	c.JSON(200, gin.H{
		"message": "Deal Assigned to Product",
	})
}

type GetProducts struct {
	// empty struct
	Id int `json:"id"`
}

func GetProductsController(c *gin.Context) {
	// return all products
	GetAllProducts := GetProducts{}
	var p *deal.Product
	c.BindJSON(&GetAllProducts)
	product := p.GetProduct(GetAllProducts.Id)

	c.JSON(200, gin.H{
		"message": product,
	})
}
