package main

import (
	"fmt"
	"project/auth"
	"project/database"
	"project/handler"
	reposerviceCampaign "project/reposervice/reposervice-campaign"
	reposerviceUser "project/reposervice/reposervice-user"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "testdb"
)

func main() {
	// Connect database to postgresql
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+" password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = database.CreateTable(db)
	if err != nil {
		panic(err)
	}

	// tahap 1 - make repository
	userRepository := reposerviceUser.NewRepository(db)
	campaignRepository := reposerviceCampaign.NewRepository(db)

	// tahap 2 - make service
	userService := reposerviceUser.NewService(userRepository)
	campaignService := reposerviceCampaign.NewService(campaignRepository)

	// tahap 4
	authService := auth.NewService()

	// tahap 3 - make user handler to service
	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	router := gin.Default()
	router.Static("/images", "./images")

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", auth.AuthMiddleware(authService, userService), userHandler.UploadAvataric)

	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)

	router.Run()
}
