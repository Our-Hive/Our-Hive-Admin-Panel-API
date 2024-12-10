package controller

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/request"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/configuration/security"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RecommendedContentController struct {
	recommendedContentHandler application.IRecommendedContentHandler
}

func NewRecommendedContentController(recommendedContentHandler application.IRecommendedContentHandler) *RecommendedContentController {
	return &RecommendedContentController{recommendedContentHandler: recommendedContentHandler}
}

func (r RecommendedContentController) InitRoutes(router *gin.Engine) {
	router.POST("/recommended-content", security.JwtMiddleware, security.AdminRoleMiddleware, r.CreateRecommendedContent)
	router.GET("/recommended-content", security.JwtMiddleware, r.GetAllRecommendedContent)
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
// @Success 200 {array} model.DigitalContent
// @Failure 500 {object} string
// @Router /recommended-content [get]
func (r RecommendedContentController) GetAllRecommendedContent(c *gin.Context) {
	content, httpStatus, err := r.recommendedContentHandler.GetAll()

	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpStatus, content)
}
