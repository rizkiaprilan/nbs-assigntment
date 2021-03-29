package helpers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
