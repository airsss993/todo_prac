package routes

import (
	"github.com/gin-gonic/gin"
	"trackit/backend/controllers"
)

func UseRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.SignUp)
}
