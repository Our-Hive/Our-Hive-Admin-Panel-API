package main

import (
	_ "github.com/Our-Hive/Our-Hive-Admin-Panel-API/cmd/docs"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/configuration"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	configuration.LoadEnvVariables()
	infrastructure.InitializeFirebase()

	router := gin.Default()
	// Enable CORS for localhost:4200
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	}))
	generationController := infrastructure.InitializeGenerationController()
	generationController.InitRoutes(router)

	imageController := infrastructure.InitializeImageController()
	imageController.InitRoutes(router)

	contactLineController := infrastructure.InitializeContactLineController()
	contactLineController.InitRoutes(router)

	recommendedContentController := infrastructure.InitializeRecommendedContentController()
	recommendedContentController.InitRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err := router.Run()

	if err != nil {
		panic(err)
	}
}
