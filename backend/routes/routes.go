package routes

import (
	"github.com/gin-gonic/gin"
	"todo_prac/backend/controllers"
	"todo_prac/backend/middleware"
)

func UseRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	auth := r.Group("/auth")
	auth.Use(middleware.AuthRequest)
	{
		auth.POST("/task", controllers.CreateTask)    // POST /auth/task
		auth.GET("/tasks", controllers.GetTasks)      // GET /auth/tasks
		auth.PUT("/tasks", controllers.UpdateTask)    // PUT /auth/tasks
		auth.DELETE("/tasks", controllers.DeleteTask) // DELETE /auth/tasks
	}
}
