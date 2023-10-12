package main

import (
	api "github.com/AdamPekny/IIS/backend"
	"github.com/AdamPekny/IIS/backend/utils"
)

func main() {
	utils.Migrate_all()

	router := api.Router()
	router.Run("0.0.0.0:8080")
}
