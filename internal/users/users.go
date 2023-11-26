package users

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/blsmxiu47/fft-api-go/internal/core/app"
)

func main() {
	// Initializing app defined in core.
	// TODO: probably move this as we branch out to more than jut users data
	a := &app.App{}
	// Load environment variables
	err := godotenv.Load("env/.env.local")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)

	a.Run(":8010")
}
