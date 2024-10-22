package routes

import (
	"backendShedulerPro/internal/database/controllers"

	"github.com/gin-gonic/gin"
)

type MailingRoutesController struct {
	mailingController controllers.MailingController
}

// Constructor 
func NewMailingController(mailingController controllers.MailingController) MailingRoutesController {
	return MailingRoutesController{mailingController}
}

func (mrc *MailingRoutesController) MailingRoute(rg *gin.RouterGroup) {
	router := rg.Group("mailings")
	router.POST("/", mrc.mailingController.CreateMailing)
	router.GET("/", mrc.mailingController.FindAllMailing)
	router.GET("/:id", mrc.mailingController.FindMailingById)
	router.PUT("/:id", mrc.mailingController.UpdateMailing)
	router.DELETE("/:id", mrc.mailingController.DeleteMailing)
}
