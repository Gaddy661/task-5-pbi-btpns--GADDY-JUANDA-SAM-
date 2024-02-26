package router

import (
	"pbi-final-task-go-api/controllers"
	"pbi-final-task-go-api/middlewares"

	"github.com/gin-gonic/gin"
)

func PhotoRoutes(r *gin.Engine) {
	photos := r.Group("/photos")
	{
		photos.POST("", middlewares.RequireAuth, controllers.AddPhoto)
		photos.GET("", middlewares.RequireAuth, controllers.GetPhotos)
		photos.PUT("/:photoId", middlewares.RequireAuth, controllers.EditPhoto)
		photos.DELETE("/:photoId", middlewares.RequireAuth, controllers.DeletePhoto)
	}
}
