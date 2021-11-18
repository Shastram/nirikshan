package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"nirikshan-backend/api/handlers"
	"nirikshan-backend/api/routes"
	"nirikshan-backend/pkg/records"
	"nirikshan-backend/pkg/services"
	"nirikshan-backend/pkg/siteconfigs"
	"nirikshan-backend/pkg/user"
	"nirikshan-backend/pkg/utils"
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

	userCollection := db.Collection(utils.UserCollection)
	userRepo := user.NewRepo(userCollection)

	siteConfigCollection := db.Collection(utils.SiteConfigCollection)
	siteRepo := siteconfigs.NewRepo(siteConfigCollection)

	userRecordCollection := db.Collection(utils.UserRecordsCollection)
	userRecordRepo := records.NewRepo(userRecordCollection)
	applicationService := services.NewService(userRepo, siteRepo, userRecordRepo, db)

	runRestServer(applicationService)
}

func runRestServer(applicationService services.ApplicationService) {
	// Starting REST server using Gin
	gin.SetMode(gin.ReleaseMode)
	gin.EnableJsonDecoderDisallowUnknownFields()
	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	app.Any("/google/*proxyPath", handlers.Proxy(applicationService))
	routes.UserRouter(app, applicationService)
	routes.SiteConfigRouter(app, applicationService)
	log.Infof("Nirikshan server is successfully ready to serve on port: %s",
		utils.Port)
	err := app.Run(utils.Port)
	if err != nil {
		log.Fatalf("Failure to start nirikshan server due to %s", err)
	}
}
