package users

import (
	"github.com/blsmxiu47/fft-api-go/internal/core/app"
	"github.com/blsmxiu47/fft-api-go/internal/utils"
)

func main() {
	// Initializing app defined in core.
	// TODO: probably move this as we branch out to more than just users data
	a := &app.App{}
	a.Initialize(
		utils.GetEnv("APP_DB_USERNAME"),
		utils.GetEnv("APP_DB_PASSWORD"),
		utils.GetEnv("APP_DB_NAME"),
	)

	a.Run(":8010")
}
