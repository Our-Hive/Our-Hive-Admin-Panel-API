package controller

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/request"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/configuration/security"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RecommendedContentController struct {
	recommendedContentHandler application.IRecommendedContentHandler
}

func NewRecommendedContentController(recommendedContentHandler application.IRecommendedContentHandler) *RecommendedContentController {
	return &RecommendedContentController{recommendedContentHandler: recommendedContentHandler}
}

func (r RecommendedContentController) InitRoutes(router *gin.Engine) {
	router.POST("/recommended-content", security.JwtMiddleware, security.AdminRoleMiddleware, r.CreateRecommendedContent)
	router.GET("/recommended-content", security.JwtMiddleware, security.AdminRoleMiddleware, r.GetAllRecommendedContent)
}

// CreateRecommendedContent godoc
// @Summary Create a new recommended content
// @Description Create a new recommended content
// @Tags Recommended Content
// @Accept json
// @Produce json
// @Param body body request.CreateDigitalContent true "Recommended Content"
// @Security ApiKeyAuth
// @Success 201
// @Failure 400 {object} string
// @Router /recommended-content [post]
func (r RecommendedContentController) CreateRecommendedContent(c *gin.Context) {
	var content request.CreateDigitalContent

	if err := c.ShouldBindJSON(&content); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	httpStatus, err := r.recommendedContentHandler.Create(&content)

	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	c.Status(httpStatus)
}

// GetAllRecommendedContent godoc
// @Summary Get all recommended content
// @Description Get all recommended content
// @Tags Recommended Content
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} []model.DigitalContent
// @Failure 500 {object} string
// @Router /recommended-content [get]
func (r RecommendedContentController) GetAllRecommendedContent(c *gin.Context) {
	pageSize, err := strconv.Atoi(c.DefaultQuery("size", "10"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startAfter := c.DefaultQuery("startAfter", "")

	contents, httpStatus, err := r.recommendedContentHandler.GetAll(pageSize, startAfter)

	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpStatus, contents)
}
