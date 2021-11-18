package routes

import (
	"github.com/gin-gonic/gin"
	"nirikshan-backend/api/handlers"
	"nirikshan-backend/pkg/services"
)

// UserRouter creates all the required routes for user authentications purposes.
func UserRouter(router *gin.Engine, service services.ApplicationService) {
	router.GET("/", handlers.Status(service))
	//router.POST("/signup", rest.UserSignUp(service))
	//router.Use(middleware.JwtMiddleware())
	//router.GET("/get_user/:uid", rest.GetUser(service))
	//router.GET("/users", rest.FetchUsers(service))
}