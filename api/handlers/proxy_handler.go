package handlers

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"nirikshan-backend/api/presenter"
	"nirikshan-backend/pkg/entities"
	"nirikshan-backend/pkg/services"
	"nirikshan-backend/pkg/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	ua "github.com/mileusna/useragent"
	log "github.com/sirupsen/logrus"
)

func Proxy(service services.ApplicationService,
	siteName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		configs, err := service.GetSiteData(siteName)

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
		log.Info("Client Version: ", userInfo.OSVersion)
		log.Info("Client Browser: ", userInfo.Name)

		dump := entities.UserRecords{
			SiteID:        configs.ID,
			SiteName:      configs.SiteName,
			Device:        userInfo.Device,
			Os:            userInfo.OS,
			Browser:       userInfo.Name,
			IP:            cip,
			Time:          time.Now(),
			IsBlackListed: false,
		}

		if err != nil {
			log.Error(err)
			c.JSON(utils.ErrorStatusCodes[utils.ErrServerError],
				presenter.CreateErrorResponse(utils.ErrServerError))
			return
		}

		if utils.Contains(configs.BlockedIP, cip) || utils.Contains(configs.
			BlockedBrowser, userInfo.Name) || utils.Contains(configs.
			BlockedOS, userInfo.OS) || utils.Contains(configs.
			BlockedOSVersion, userInfo.OSVersion) || utils.Contains(
			configs.BlockedLocations, c.Param("proxyPath")) || utils.
			Contains(configs.BlockedDevice, userInfo.Device) {
			log.Info(utils.Contains(configs.BlockedIP, cip),
				utils.Contains(configs.
					BlockedBrowser, userInfo.Name), utils.Contains(configs.
					BlockedOS, userInfo.OS), utils.Contains(configs.
					BlockedOSVersion, userInfo.OSVersion), utils.Contains(
					configs.BlockedLocations, c.Param("proxyPath")), utils.
					Contains(configs.BlockedDevice, userInfo.Device))
			dump.IsBlackListed = true
			err = service.CreateDump(&dump)
			log.Warn("blacklisted entry")
			c.JSON(utils.ErrorStatusCodes[utils.ErrNotAllowed],
				presenter.CreateErrorResponse(utils.ErrNotAllowed))
			return
		}
		err = ddosCounter(service, cip)
		if err != nil {
			log.Error(err)
			if err == utils.ErrNoticeBan {
				err := service.SendMessage(utils.DDoSTemplate(cip,
					userInfo.Name, userInfo.OS))
				if err != nil {
					log.Error(err)
					return
				}
			}
			dump.IsBlackListed = true
			err = service.CreateDump(&dump)
			c.JSON(utils.ErrorStatusCodes[utils.ErrNotAllowed],
				presenter.CreateErrorResponse(utils.ErrNotAllowed))
			return
		}
		err = service.CreateDump(&dump)
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

func ddosCounter(service services.ApplicationService, ip string) error {
	getIpCount, err := service.GetKey(ip)
	if err == redis.Nil || getIpCount == "" {
		err = service.PutData(ip, strconv.Itoa(0))
		if err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}

	count, err := strconv.Atoi(getIpCount)
	if err != nil {
		return err
	}
	if count > utils.DdosCountLimit {
		return utils.ErrDos
	}
	err = service.PutData(ip, strconv.Itoa(count+1))
	if err != nil {
		return err
	}
	if count == utils.DdosCountLimit {
		return utils.ErrNoticeBan
	}
	return nil
}
