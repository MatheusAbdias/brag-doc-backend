package routers

import (
	controllers "github.com/MatheusAbdias/brag-doc-backend/internal/controllers/events"
	"github.com/gin-gonic/gin"
)

func RegisterEventRouters(router *gin.RouterGroup, eventController controllers.Controller) {
	eventRouter := router.Group("events")

	eventRouter.POST("", eventController.CreateEvent)
	eventRouter.GET(":id", eventController.GetEvent)
	eventRouter.GET("", eventController.ListEvents)
	eventRouter.PATCH(":id", eventController.UpdateEvent)
	eventRouter.DELETE(":id", eventController.DeleteEvent)
}
