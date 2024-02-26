package router

import (
	"pbi-final-task-go-api/controllers"
	"pbi-final-task-go-api/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.POST("/register", controllers.Register)
		users.POST("/login", controllers.Login)
		users.PUT("/:userId", middlewares.RequireAuth, controllers.Update)
		users.DELETE("/:userId", middlewares.RequireAuth, controllers.Delete)
	}
}
