package main

import (
	_ "github.com/Our-Hive/Our-Hive-Admin-Panel-API/cmd/docs"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/configuration"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	configuration.LoadEnvVariables()
	infrastructure.InitializeFirebase()

	router := gin.Default()

	generationController := infrastructure.InitializeGenerationController()
	generationController.InitRoutes(router)

	imageController := infrastructure.InitializeImageController()
	imageController.InitRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err := router.Run()

	if err != nil {
		panic(err)
	}
}
