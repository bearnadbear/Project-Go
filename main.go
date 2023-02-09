package main

import (
	"fmt"
	"project/database"
	"project/handler"
	"project/reserv"

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

	// Struct User / Table Users
	// var users []user.User

	// Find struct User to Table Users
	// db.Find(&users)

	// tahap 1 - make repository
	userRepository := reserv.NewRepository(db)

	// tahap 2 - make service
	userService := reserv.NewService(userRepository)

	// tahap 3 - make user handler to service
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)

	router.Run()
}
