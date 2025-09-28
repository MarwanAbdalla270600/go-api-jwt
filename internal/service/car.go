package service

import (
	"go-api-jwt/internal/entity"

	"github.com/jmoiron/sqlx"
)

type CarInterface interface {
	GetCars() ([]entity.Car, error)
	GetCarById(id string) (*entity.Car, error)
}

type carService struct {
	db *sqlx.DB
}

func NewCarService(db *sqlx.DB) CarInterface {
	return &carService{db: db}
}

func (s *carService) GetCars() ([]entity.Car, error) {
	var cars []entity.Car
	err := s.db.Select(&cars, "SELECT * FROM cars")
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func (s *carService) GetCarById(id string) (*entity.Car, error) {
	var car entity.Car
	err := s.db.Get(&car, "SELECT * FROM cars WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &car, nil
}
