package service

import (
	"go-api-jwt/internal/entity"
	"go-api-jwt/internal/utils"

	"github.com/jmoiron/sqlx"
)

type AuthServiceInterface interface {
	Register(data entity.RegisterRequest) (*entity.UserDTO, error)
	Login(data entity.LoginRequest) error
}

type authService struct {
	db *sqlx.DB
}

func NewAuthService(db *sqlx.DB) AuthServiceInterface {
	return &authService{db: db}
}

func (s *authService) Register(data entity.RegisterRequest) (*entity.UserDTO, error) {
	uuid := utils.GenerateUUID()
	_, err := s.db.Exec(`INSERT INTO users (id, first_name, last_name, email, hashed_password)
		VALUES (?, ?, ?, ?, ?)`,
		uuid, data.FirstName, data.LastName, data.Email, utils.HashPassword(data.Password),
	)

	if err != nil {
		return nil, err
	}

	return &entity.UserDTO{Id: uuid,
			FirstName: data.FirstName,
			LastName:  data.LastName,
			Email:     data.Email},
		nil
}

func (s *authService) Login(data entity.LoginRequest) error {
	return nil
}
