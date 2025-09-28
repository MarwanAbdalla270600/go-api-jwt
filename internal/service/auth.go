package service

import (
	"errors"
	"go-api-jwt/internal/entity"
	"go-api-jwt/internal/utils"

	"github.com/jmoiron/sqlx"
)

type AuthServiceInterface interface {
	Register(data entity.RegisterRequest) (*entity.UserDTO, error)
	Login(data entity.LoginRequest) (*entity.LoginResponse, error)
}

type authService struct {
	db *sqlx.DB
}

func NewAuthService(db *sqlx.DB) AuthServiceInterface {
	return &authService{db: db}
}

func (s *authService) Register(data entity.RegisterRequest) (*entity.UserDTO, error) {
	var user entity.User
	err := s.db.Get(&user, `
    INSERT INTO users (first_name, last_name, email, password)
    VALUES ($1, $2, $3, $4)
    RETURNING id, first_name, last_name, email, role, created_at, updated_at
`,
		data.FirstName, data.LastName, data.Email, utils.HashPassword(data.Password),
	)
	if err != nil {
		return nil, err
	}

	return &entity.UserDTO{Id: user.Id,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		nil
}

func (s *authService) Login(data entity.LoginRequest) (*entity.LoginResponse, error) {
	//get user from database
	var user entity.User
	err := s.db.Get(&user, `SELECT id, first_name, last_name, email, password, role FROM users 
	WHERE email = $1`, data.Email)
	if err != nil {
		return nil, err
	}

	//check password match
	if !utils.ComparePasswords(user.HashedPassword, data.Password) {
		return nil, errors.New("invalid password")
	}

	//create response with token
	response := entity.LoginResponse{
		Token: utils.GenerateJWT(user.Id, user.Email, user.Role),
		User: entity.UserDTO{
			Id:        user.Id,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      user.Role,
		},
	}

	return &response, nil
}
