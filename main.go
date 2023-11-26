package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// App struct defined in app.go
	a := App{}
	// Load environment variables
	err := godotenv.Load("env/.env.local")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// This does __nothing so far__
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)

	a.Run(":8010")
}
