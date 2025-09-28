package service

import (
	"go-api-jwt/internal/entity"

	"github.com/jmoiron/sqlx"
)

type CarInterface interface {
	GetCars() ([]entity.Car, error)
	GetCarById() (*entity.Car, error)
}

type carService struct {
	db *sqlx.DB
}

func NewCarService(db *sqlx.DB) CarInterface {
	return &carService{db: db}
}

func (s *carService) GetCars() ([]entity.Car, error) {
	return nil, nil
}

func (s *carService) GetCarById() (*entity.Car, error) {
	return nil, nil

}
