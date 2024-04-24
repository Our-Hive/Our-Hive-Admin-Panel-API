package controller

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ImageController struct {
	imageHandler handler.IImageHandler
}

func (i ImageController) InitRoutes(router *gin.Engine) {
	router.GET("/images", i.GetAll)
}

func NewImageController(imageHandler handler.IImageHandler) *ImageController {
	return &ImageController{imageHandler: imageHandler}
}

func (i ImageController) GetAll(c *gin.Context) {
	pageSize, err := strconv.Atoi(c.DefaultQuery("size", "10"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startAfter := c.DefaultQuery("startAfter", "")

	images, httpStatus, err := i.imageHandler.GetAll(pageSize, startAfter)

	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpStatus, images)
}
