package controller

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/request"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ContactLineController struct {
	contactLineHandler handler.IContactLineHandler
}

func NewContactLineController(contactLineHandler handler.IContactLineHandler) *ContactLineController {
	return &ContactLineController{contactLineHandler: contactLineHandler}
}

func (controller ContactLineController) InitRoutes(router *gin.Engine) {
	//router.POST("/contact-line", security.JwtMiddleware, security.AdminRoleMiddleware, controller.CreateContactLine)
	router.POST("/contact-line", controller.CreateContactLine)
	//router.GET("/contact-line", security.JwtMiddleware, security.AdminRoleMiddleware, controller.GetAll)
	router.GET("/contact-line", controller.GetAll)
}

// CreateContactLine godoc
// @Summary Create a contact line
// @Description Create a contact line
// @Tags Contact Line
// @Accept json
// @Produce json
// @Param body body request.CreateContactLine true "Contact Line"
// @Success 201
// @Failure 400 {object} string
// @Router /contact-line [post]
func (controller ContactLineController) CreateContactLine(c *gin.Context) {
	var contactLine request.CreateContactLine

	if err := c.ShouldBindJSON(&contactLine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	httpStatus, err := controller.contactLineHandler.Create(&contactLine)

	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	c.Status(httpStatus)
}

// GetAll godoc
// @Summary Get all contact lines
// @Description Get all contact lines with pagination
// @Tags Contact Line
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param size query int false "Page size"
// @Param startAfter query string false "Start after"
// @Success 200 {array} model.ContactLine "Success"
// @Failure 404
// @Router /contact-line [get]
func (controller ContactLineController) GetAll(c *gin.Context) {
	pageSize, err := strconv.Atoi(c.DefaultQuery("size", "10"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startAfter := c.DefaultQuery("startAfter", "")

	contactLines, httpStatus, err := controller.contactLineHandler.GetAll(pageSize, startAfter)

	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpStatus, contactLines)
}
