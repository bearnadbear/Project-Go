package main

import (
	"fmt"
	"project/auth"
	"project/database"
	"project/handler"
	"project/reserv_campaign"
	"project/reserv_user"

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
	userRepository := reserv_user.NewRepository(db)
	campaignRepository := reserv_campaign.NewRepository(db)

	// tahap 2 - make service
	userService := reserv_user.NewService(userRepository)
	campaignService := reserv_campaign.NewService(campaignRepository)

	campaign, _ := campaignService.FindCampaign(4)
	fmt.Println(len(campaign))
	// tahap 4
	authService := auth.NewService()

	// tahap 3 - make user handler to service
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", auth.AuthMiddleware(authService, userService), userHandler.UploadAvataric)

	router.Run()
}
