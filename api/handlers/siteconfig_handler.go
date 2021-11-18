package handlers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"nirikshan-backend/api/presenter"
	"nirikshan-backend/pkg/entities"
	"nirikshan-backend/pkg/services"
	"nirikshan-backend/pkg/utils"
)

func CreateSiteConfig(service services.ApplicationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request entities.SiteConfigs
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(utils.ErrorStatusCodes[utils.ErrInvalidRequest],
				presenter.CreateErrorResponse(utils.ErrInvalidRequest))
			return
		}
		err = service.CreateSiteData(&request)

		if err != nil {
			c.JSON(utils.ErrorStatusCodes[utils.ErrServerError],
				presenter.CreateErrorResponse(utils.ErrServerError))
			return
		}

		c.JSON(http.StatusOK, presenter.CreateSuccessResponse(
			request,
			"site config has been created!"))
	}
}

func GetSiteConfig(service services.ApplicationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		site := c.Query("site")
		if site == "" {
			c.JSON(utils.ErrorStatusCodes[utils.ErrInvalidRequest],
				presenter.CreateErrorResponse(utils.ErrInvalidRequest))
			return
		}
		configs, err := service.GetSiteData(site)
		if err != nil {
			log.Error(err)
			if err == mongo.ErrNoDocuments {
				c.JSON(utils.ErrorStatusCodes[utils.ErrInvalidSite],
					presenter.CreateErrorResponse(utils.ErrInvalidSite))
				return
			}
			c.JSON(utils.ErrorStatusCodes[utils.ErrServerError],
				presenter.CreateErrorResponse(utils.ErrServerError))
			return
		}
		c.JSON(http.StatusOK, presenter.CreateSuccessResponse(
			configs,
			"site config fetched!"))
	}
}

func GetSiteDump(service services.ApplicationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		site := c.Query("site")
		if site == "" {
			c.JSON(utils.ErrorStatusCodes[utils.ErrInvalidRequest],
				presenter.CreateErrorResponse(utils.ErrInvalidRequest))
			return
		}
		dump, err := service.GetDump(site)
		if err != nil {
			log.Error(err)
			if err == mongo.ErrNoDocuments {
				c.JSON(utils.ErrorStatusCodes[utils.ErrInvalidSite],
					presenter.CreateErrorResponse(utils.ErrInvalidSite))
				return
			}
			c.JSON(utils.ErrorStatusCodes[utils.ErrServerError],
				presenter.CreateErrorResponse(utils.ErrServerError))
			return
		}
		var counter = 0
		for _, data := range *dump {
			if data.IsBlackListed {
				counter++
			}
		}
		c.JSON(http.StatusOK, presenter.CreateSuccessResponse(
			entities.UserRecordsResponse{
				Logs:            dump,
				TotalLength:     len(*dump),
				BlacklistLength: counter,
			},
			"site dump fetched!"))
	}
}
