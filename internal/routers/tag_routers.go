package routers

import (
	controllers "github.com/MatheusAbdias/brag-doc-backend/internal/controllers/tags"
	"github.com/gin-gonic/gin"
)

func RegisterTagRouters(router *gin.RouterGroup, Controller controllers.Controller) {
	tagsRouter := router.Group("tags")

	tagsRouter.POST("", Controller.CreateTag)
	tagsRouter.GET(":id", Controller.GetTag)
	tagsRouter.GET("", Controller.ListTags)
	tagsRouter.PATCH(":id", Controller.UpdateTag)
	tagsRouter.DELETE(":id", Controller.DeleteTag)
}
