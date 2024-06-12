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
	router.POST("/recommended-content", security.JwtMiddleware, security.AdminRoleMiddleware, controller.CreateRecommendedContent)
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
