package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"nirikshan-backend/api/handlers"
	"nirikshan-backend/api/routes"
	"nirikshan-backend/pkg/entities"
	"nirikshan-backend/pkg/records"
	"nirikshan-backend/pkg/services"
	"nirikshan-backend/pkg/siteconfigs"
	"nirikshan-backend/pkg/user"
	"nirikshan-backend/pkg/utils"
	"time"
)

func main() {
	log.Info("Nirikshan is booting")

	db, err := utils.DatabaseConnection()
	if err != nil {
		log.Fatal("database connection error $s", err)
	}

	err = utils.CreateCollection(utils.UserCollection, db)
	if err != nil {
		log.Fatalf("failed to create collection  %s", err)
	}

	err = utils.CreateCollection(utils.UserRecordsCollection, db)
	if err != nil {
		log.Fatalf("failed to create collection  %s", err)
	}

	err = utils.CreateIndex(utils.UserCollection, utils.UsernameField, db)
	if err != nil {
		log.Fatalf("failed to create index  %s", err)
	}

	err = utils.CreateCollection(utils.SiteConfigCollection, db)
	if err != nil {
		log.Fatalf("failed to create collection  %s", err)
	}

	err = utils.CreateIndex(utils.SiteConfigCollection, utils.SiteNameField, db)
	if err != nil {
		log.Fatalf("failed to create index  %s", err)
	}

	userCollection := db.Collection(utils.UserCollection)
	userRepo := user.NewRepo(userCollection)

	siteConfigCollection := db.Collection(utils.SiteConfigCollection)
	siteRepo := siteconfigs.NewRepo(siteConfigCollection)

	userRecordCollection := db.Collection(utils.UserRecordsCollection)
	userRecordRepo := records.NewRepo(userRecordCollection)

	rdb := utils.RedisDatabaseConnection()
	teleBot := utils.InitialiseTelegramBot()
	applicationService := services.NewService(userRepo, siteRepo,
		userRecordRepo, db, rdb, teleBot)
	securityDefinitions := setupDatabase(applicationService)
	runRestServer(applicationService, securityDefinitions)
}

func setupDatabase(service services.ApplicationService) *utils.
	SecurityPolicyDefinition {
	var c utils.SecurityPolicyDefinition
	conf, err := c.GetConf()
	if err != nil {
		log.Warnf("No security policies loaded due to %s", err)
	}
	log.Warn("Security Policies found!")
	for _, siteData := range conf.SiteConfigs {
		configs := entities.SiteConfigs{
			SiteName:         siteData.SiteData.SiteName,
			ForwardingURL:    siteData.SiteData.ForwardingURL,
			BlockedOS:        siteData.SiteData.BlockedOs,
			BlockedBrowser:   siteData.SiteData.BlockedBrowser,
			BlockedDevice:    siteData.SiteData.BlockedDevice,
			BlockedOSVersion: siteData.SiteData.BlockedOSVersion,
			BlockedLocations: siteData.SiteData.BlockedLocations,
			BlockedIP:        siteData.SiteData.BlockedIPs,
			CreatedAt:        time.Now(),
			UpdatedAt:        "",
		}
		err := service.CreateSiteData(&configs)
		if err != nil {
			if err == utils.ErrUserExists {
				log.Warnf("Policy for %s already exists", configs.SiteName)
				continue
			}
			log.Errorf("Unable to create security policy entry due to %s", err)
		}
	}
	return &c
}

func runRestServer(applicationService services.ApplicationService,
	definition *utils.SecurityPolicyDefinition) {
	// Starting REST server using Gin
	gin.SetMode(gin.ReleaseMode)
	gin.EnableJsonDecoderDisallowUnknownFields()
	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	if definition.NirikshanVersion != "" {
		for _, siteData := range definition.SiteConfigs {
			app.Any(fmt.Sprintf("/%s/*proxyPath", siteData.SiteData.SiteName),
				handlers.Proxy(applicationService,
					siteData.SiteData.SiteName))
		}
	}
	routes.UserRouter(app, applicationService)
	routes.SiteConfigRouter(app, applicationService)
	log.Infof("Nirikshan server is successfully ready to serve on port: %s",
		utils.Port)
	err := app.Run(utils.Port)
	if err != nil {
		log.Fatalf("Failure to start nirikshan server due to %s", err)
	}
}
