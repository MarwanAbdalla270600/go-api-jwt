package controller

import (
	"go-api-jwt/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type carController struct {
	service service.CarInterface
}

func NewCarController(service service.CarInterface) *carController {
	return &carController{service: service}
}

func (c *carController) GetCars(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "all cars")
}

func (c *carController) GetCarById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "car by id")
}
