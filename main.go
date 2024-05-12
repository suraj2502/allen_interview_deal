package main

import (
	"log"
	"suraj_projects/allen_interview/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	ENVIRONMENT string
	PORT        string
)

func setEnvironment() {

	PORT = "8080"

	log.Println("Port: " + PORT)
}
func initialize() {
	setEnvironment()
	// // connector.InitializeConnection()
	// // config.DoInit(ENVIRONMENT)
	// // logger.InitApiLogger(ENVIRONMENT)
	// // validator.InitValidator()

	// // backend_services.InitializeConnection()
	// // connector.InitializeConnection()
	// // middleware.CreateBucket()

	// // // Initializing constants file
	// // constants.InitializeConstants()
	// // populateChannelManagerDBData()
	// // populateHotelCodesForMMTProduct()

	// // workflow.Initialize(config.Config.GetString("logFile"))
	// workflow.Initialize("/opt/logs/human_being.json")
	// // go metrics.InitMetrics()
	// // metrics.InitStatsReporter()

}

func main() {
	initialize()
	gin.SetMode(gin.ReleaseMode)
	mainRouter := gin.New()
	// Apply CORS middleware to allow requests from any origin with the specified headers and methods
	mainRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow requests from any origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	routers.InitRoutes(mainRouter)
	mainRouter.Run(":" + PORT)
}
