package utils

import (
	"fmt"
	"os"

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

func GenerateUUID() string {
	return uuid.New().String()
}
