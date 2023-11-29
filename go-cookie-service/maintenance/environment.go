package maintenance

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// LoadEnv Получение .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// GetEnv Получение переменных из окружения
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
