package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nirikshan-backend/api/presenter"
	"nirikshan-backend/pkg/services"
)

func Status(service services.ApplicationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, presenter.CreateSuccessResponse(
			"URL: https://github."+
				"com/Shastram/nirikshan-backend",
			"Nirikshan Backend is healthy"))
	}
}
