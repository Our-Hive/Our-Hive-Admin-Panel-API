package main

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/configuration"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure"
	"github.com/gin-gonic/gin"
)

func main() {
	configuration.LoadEnvVariables()
	infrastructure.InitializeFirebase()

	router := gin.Default()

	generationController := infrastructure.InitializeGenerationController()
	generationController.InitRoutes(router)

	err := router.Run()

	if err != nil {
		panic(err)
	}
}
