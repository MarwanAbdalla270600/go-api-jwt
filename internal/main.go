package main

import (
	"go-api-jwt/internal/controller"
	"go-api-jwt/internal/service"
	"go-api-jwt/internal/utils"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	//load env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, falling back to system env")
	}

	//connect to database
	dsn := utils.BuildDatabaseUrl()
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln("failed to connect:", err)
	}
	defer db.Close()

	//initialize singletons
	authService := service.NewAuthService(db)
	authController := controller.NewAuthController(authService)

	router.POST("/auth/register", authController.Register)
	router.POST("/auth/login", authController.Login)
	router.Run(":8080")
}
