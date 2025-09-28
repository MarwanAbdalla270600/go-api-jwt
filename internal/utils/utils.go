package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func BuildDatabaseUrl() string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database)
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

func ComparePasswords(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func GenerateUUID() string {
	return uuid.New().String()
}

func GenerateJWT(userID string, email string) string {
	claims := jwt.MapClaims{
		"sub":   userID,                               // subject (user identifier)
		"email": email,                                // custom claim
		"exp":   time.Now().Add(time.Hour * 1).Unix(), // expires in 1 hour
		"iat":   time.Now().Unix(),                    // issued at
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := os.Getenv("SECRET")
	tokenString, _ := token.SignedString(jwtSecret)

	return tokenString
}
