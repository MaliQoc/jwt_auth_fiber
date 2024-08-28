package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// LoadConfig yükler .env dosyasını ve verilen anahtarın değerini döner
func LoadConfig(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv(key)
}

