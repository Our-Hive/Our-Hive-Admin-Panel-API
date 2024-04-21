package controller

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/request"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GenerationController struct {
	generationHandler handler.IGenerationHandler
}

func (g GenerationController) InitRoutes(router *gin.Engine) {
	router.POST("/generation", g.GenerateIAImage)
}

func NewGenerationController(generationHandler handler.IGenerationHandler) *GenerationController {
	return &GenerationController{generationHandler: generationHandler}
}

func (g GenerationController) GenerateIAImage(c *gin.Context) {
	var body request.GenerateIAImage
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, httpStatus, err := g.generationHandler.GenerateIAImage(&body)

	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpStatus, response)
}
