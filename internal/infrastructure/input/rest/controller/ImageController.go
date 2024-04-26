package controller

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/handler"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"strconv"
)

type ImageController struct {
	imageHandler  handler.IImageHandler
	uploadHandler handler.IUploadHandler
}

func NewImageController(imageHandler handler.IImageHandler, uploadHandler handler.IUploadHandler) *ImageController {
	return &ImageController{imageHandler: imageHandler, uploadHandler: uploadHandler}
}

func (i ImageController) InitRoutes(router *gin.Engine) {
	router.GET("/images", i.GetAll)
	router.POST("/images", i.Upload)
}

// GetAll godoc
// @Summary Get all images
// @Description Get all images with pagination
// @Tags images
// @Accept  json
// @Produce  json
// @Param size query int false "Page size"
// @Param startAfter query string false "Start after"
// @Success 200 {array} response.Image "Success"
// @Failure 404
// @Router /images [get]
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

// Upload godoc
// @Summary Upload an image
// @Description Upload an image
// @Tags images
// @Accept  multipart/form-data
// @Produce  json
// @Param image formData file true "Image"
// @Success 200 {object} response.UploadImage "Success"
// @Failure 400
// @Router /images [post]
func (i ImageController) Upload(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileHeader, err := file.Open()

	contentType := file.Header.Get("Content-Type")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer func(fileHeader multipart.File) {
		err = fileHeader.Close()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}(fileHeader)

	fileData := make([]byte, file.Size)
	_, err = fileHeader.Read(fileData)

	resp, httpStatus, err := i.uploadHandler.Upload(fileData, file.Filename, contentType)

	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpStatus, resp)
}
