package main

import (
	"go-api-jwt/internal/controller"
	"go-api-jwt/internal/middleware"
	"go-api-jwt/internal/service"
	"go-api-jwt/internal/utils"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // postgres driver
)

func main() {
	router := gin.Default()

	//load env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, falling back to system env")
	}

	router.Use(cors.Default())


	//connect to database
	dsn := utils.BuildDatabaseUrl()
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln("failed to connect:", err)
	}
	defer db.Close()

	//initialize singletons
	authService := service.NewAuthService(db)
	authController := controller.NewAuthController(authService)

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "<h1>Enta GAYYY</h1>")
	})

	router.POST("/auth/register", authController.Register)
	router.POST("/auth/login", authController.Login)

	//testroute
	router.GET("/protected", middleware.AuthMiddleware(), func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, "YOU DID IT!!")
	})
	router.Run(":8080")
}
