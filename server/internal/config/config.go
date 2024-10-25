package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func LoadConfig(){
	cwd, _ := os.Getwd()
    log.Println("Current working directory:", cwd)

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
}
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}