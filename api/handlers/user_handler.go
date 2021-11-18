package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nirikshan-backend/api/presenter"
	"nirikshan-backend/pkg/entities"
	"nirikshan-backend/pkg/services"
	"nirikshan-backend/pkg/utils"
)

func Status(service services.ApplicationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, presenter.CreateSuccessResponse(
			"URL: https://github."+
				"com/Shastram/nirikshan-backend",
			"Nirikshan Backend is healthy"))
	}
}

func SignUp(service services.ApplicationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request entities.User
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(utils.ErrorStatusCodes[utils.ErrInvalidRequest],
				presenter.CreateErrorResponse(utils.ErrInvalidRequest))
			return
		}
		err = service.CreateUser(&request, c)

		if err != nil {
			c.JSON(utils.ErrorStatusCodes[utils.ErrServerError],
				presenter.CreateErrorResponse(utils.ErrServerError))
			return
		}

		c.JSON(http.StatusOK, presenter.CreateSuccessResponse("User created!",
			"User has been created successfully!"))
	}
}

func Login(service services.ApplicationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request entities.User
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(utils.ErrorStatusCodes[utils.ErrInvalidRequest],
				presenter.CreateErrorResponse(utils.ErrInvalidRequest))
			return
		}
		if request.UserName == "" || request.Password == "" {
			c.JSON(utils.ErrorStatusCodes[utils.ErrInvalidRequest],
				presenter.CreateErrorResponse(utils.ErrInvalidRequest))
			return
		}
		user, err := service.FindUserByUsername(request.UserName)
		if err != nil {
			c.JSON(utils.ErrorStatusCodes[utils.ErrServerError],
				presenter.CreateErrorResponse(utils.ErrServerError))
			return
		}
		err = service.CheckPasswordHash(user.Password, request.Password)
		if err != nil {
			c.JSON(utils.ErrorStatusCodes[utils.ErrUnauthorized],
				presenter.CreateErrorResponse(utils.ErrUnauthorized))
			return
		}
		token, err := user.GetSignedJWT()
		c.JSON(http.StatusOK, presenter.CreateSuccessResponse(entities.UserResponse{
			JwtToken: token,
			User:     user,
		},
			"User logged in successfully!"))
	}
}
