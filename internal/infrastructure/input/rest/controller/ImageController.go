package controller

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/configuration/security"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"strconv"
)

type ImageController struct {
	imageHandler  application.IImageHandler
	uploadHandler application.IUploadHandler
}

func NewImageController(imageHandler application.IImageHandler, uploadHandler application.IUploadHandler) *ImageController {
	return &ImageController{imageHandler: imageHandler, uploadHandler: uploadHandler}
}

func (i ImageController) InitRoutes(router *gin.Engine) {
	router.GET("/images", security.JwtMiddleware, security.AdminRoleMiddleware, i.GetAll)
	router.GET("/images/approval", security.JwtMiddleware, i.GetByApprovalStatus)
	router.POST("/images", security.JwtMiddleware, security.AdminRoleMiddleware, i.Upload)
	router.PUT("/images/:id", security.JwtMiddleware, security.AdminRoleMiddleware, i.Approve)
}

// GetAll godoc
// @Summary Get all images
// @Description Get all images with pagination
// @Tags images
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
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

// GetByApprovalStatus godoc
// @Summary Get all images by approved status
// @Description Get all images by approved status
// @Tags images
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param approved query bool false "Approved"
// @Param size query int false "Page size"
// @Param startAfter query string false "Start after"
// @Success 200 {array} response.Image "Success"
// @Failure 404
// @Router /images/approval [get]
func (i ImageController) GetByApprovalStatus(c *gin.Context) {
	approved, err := strconv.ParseBool(c.DefaultQuery("approved", "false"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("size", "10"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startAfter := c.DefaultQuery("startAfter", "")

	images, httpStatus, err := i.imageHandler.GetByApprovedStatus(approved, pageSize, startAfter)

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
// @Security ApiKeyAuth
// @Param image formData file true "File"
// @Success 200 {object} response.UploadImage "Success"
// @Failure 400
// @Router /images [post]
func (i ImageController) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
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

// Approve godoc
// @Summary Approve an image
// @Description Approve an image
// @Tags images
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "Image ID"
// @Success 200
// @Failure 400
// @Router /images/{id} [put]
func (i ImageController) Approve(c *gin.Context) {
	id := c.Param("id")

	httpStatus, err := i.imageHandler.Approve(id)

	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
