package controller

import (
	"go-api-jwt/internal/entity"
	"go-api-jwt/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authController struct {
	service service.AuthServiceInterface
}

func NewAuthController(service service.AuthServiceInterface) *authController {
	return &authController{service: service}
}

func (c *authController) Register(ctx *gin.Context) {
	var registerData entity.RegisterRequest

	//bind body data of request with struct
	if err := ctx.ShouldBindJSON(&registerData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "invalid registration data",
		})
		return
	}

	userDTO, err := c.service.Register(registerData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "successfully registered user",
		"data":    userDTO,
	})
}

func (c *authController) Login(ctx *gin.Context) {
	//parsing data
	var body entity.LoginRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, "This is login endpoint")
}
