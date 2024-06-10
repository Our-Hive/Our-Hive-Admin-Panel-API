package controller

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/request"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ContactLineController struct {
	contactLineHandler handler.IContactLineHandler
}

func NewContactLineController(contactLineHandler handler.IContactLineHandler) *ContactLineController {
	return &ContactLineController{contactLineHandler: contactLineHandler}
}

func (controller ContactLineController) InitRoutes(router *gin.Engine) {
	router.POST("/contact-line", controller.CreateContactLine)
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

	c.JSON(httpStatus, nil)
}
