package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	// TODO: this path seems like bad practice
	err := godotenv.Load("../../env/.env.local")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("environment var %s not set\n", key)
	}

	return val
}
