package routes

import (
	"nirikshan-backend/api/handlers"
	"nirikshan-backend/pkg/services"

	"github.com/gin-gonic/gin"
)

// UserRouter creates all the required routes for user authentications purposes.
func UserRouter(router *gin.Engine, service services.ApplicationService) {
	router.GET("/status", handlers.Status(service))
	router.POST("/signup", handlers.SignUp(service))
	router.POST("/login", handlers.Login(service))
}
