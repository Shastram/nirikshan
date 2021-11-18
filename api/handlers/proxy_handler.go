package handlers

import (
	"github.com/gin-gonic/gin"
	ua "github.com/mileusna/useragent"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/httputil"
	"net/url"
	"nirikshan-backend/api/presenter"
	"nirikshan-backend/pkg/services"
	"nirikshan-backend/pkg/utils"
)

func Proxy(service services.ApplicationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		configs, err := service.GetSiteData("google")

		if err != nil {
			log.Error(err)
			c.JSON(utils.ErrorStatusCodes[utils.ErrServerError],
				presenter.CreateErrorResponse(utils.ErrServerError))
			return
		}
		if len(configs.BlockedIP) == 0 {
			configs.BlockedIP = []string{"0.0.0.0"}
		}

		cip := c.ClientIP()
		userAgent := c.GetHeader("User-Agent")
		userInfo := ua.Parse(userAgent)
		log.Info("Client IP: ", cip)
		log.Info("Client OS: ", userInfo.OS)
		log.Info("Client Device: ", userInfo.Device)
		log.Info("Client Device: ", userInfo.OSVersion)
		log.Info("Client Browser: ", userInfo.Name)

		if utils.Contains(configs.BlockedIP, cip) || userInfo.OS == configs.
			BlockedOS || userInfo.Device == configs.
			BlockedDevice || userInfo.Version == configs.
			BlockedOSVersion || userInfo.Name == configs.BlockedBrowser {
			c.JSON(utils.ErrorStatusCodes[utils.ErrNotAllowed],
				presenter.CreateErrorResponse(utils.ErrNotAllowed))
			return
		}

		remote, err := url.Parse(configs.ForwardingURL)
		if err != nil {
			log.Error(err)
			c.JSON(utils.ErrorStatusCodes[utils.ErrServerError],
				presenter.CreateErrorResponse(utils.ErrServerError))
			return
		}
		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			req.URL.Path = c.Param("proxyPath")
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
