package routes

import (
	"github.com/gin-gonic/gin"
	"trackit/backend/controllers"
	"trackit/middleware"
)

func UseRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", middleware.AuthRequest, controllers.Login)
	r.POST("/task", controllers.CreateTask)
	r.GET("/tasks", controllers.GetTasks)
	r.POST("/tasks", controllers.UpdateTask)
}
