package routes

import (
	"nirikshan-backend/api/handlers"
	"nirikshan-backend/api/middleware"
	"nirikshan-backend/pkg/services"

	"github.com/gin-gonic/gin"
)

// SiteConfigRouter creates all the required routes for site config purposes.
func SiteConfigRouter(router *gin.Engine, service services.ApplicationService) {
	router.Use(middleware.JwtMiddleware())
	router.GET("/site", handlers.GetSiteConfig(service))
	router.POST("/site", handlers.CreateSiteConfig(service))
	router.GET("/dump", handlers.GetSiteDump(service))
}
