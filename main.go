package main

import (
	"fmt"
	"project/auth"
	"project/database"
	"project/handler"
	sourceCampaign "project/source_campaign"
	sourceTransaction "project/source_transaction"
	sourceUser "project/source_user"

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
	userRepository := sourceUser.NewRepository(db)
	campaignRepository := sourceCampaign.NewRepository(db)
	transactionsRepository := sourceTransaction.NewRepository(db)

	// tahap 2 - make service
	userService := sourceUser.NewService(userRepository)
	campaignService := sourceCampaign.NewService(campaignRepository)
	transactionsService := sourceTransaction.NewService(transactionsRepository, campaignRepository)

	// tahap 4
	authService := auth.NewService()

	// tahap 3 - make user handler to service
	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionsService)

	router := gin.Default()
	router.Static("/images", "./images")

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", auth.AuthMiddleware(authService, userService), userHandler.UploadAvataric)

	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", auth.AuthMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", auth.AuthMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", auth.AuthMiddleware(authService, userService), campaignHandler.UploadImage)

	api.GET("/campaigns/:id/transactions", auth.AuthMiddleware(authService, userService), transactionHandler.GetCampaignTransaction)

	router.Run()
}
