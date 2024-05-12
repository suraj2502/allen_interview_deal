package routers

import (
	apiController "suraj_projects/allen_interview/api"

	"github.com/gin-gonic/gin"
)

// Other than health monitor api, all request will be processed by middlewares
func InitRoutes(mainRouter *gin.Engine) {

	api := mainRouter.Group("/api")

	api.POST("/user/create", apiController.CreateNewUserController)
	api.POST("/deal/create", apiController.CreateDealController)
	api.POST("/product/create", apiController.CreateNewProductController)
	api.POST("/user/claimdeal", apiController.ClaimDealController)
	api.POST("/user/buyproduct", apiController.ClaimDealController)
	api.POST("/product/assigndeal", apiController.AssignDealToProductController)
	api.POST("/deal/update", apiController.UpdateDealController)
	api.POST("/deal/get", apiController.GetDealController)
	api.POST("/product/get", apiController.GetProductsController)

	// api.GET("/user/getall", api.GetAllUsersController)
	// api.GET("/product/getall", api.GetAllProductsController)

}
