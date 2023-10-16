package main

import (
	api "github.com/AdamPekny/IIS/backend"
)

func main() {
	api.Migrate_all()

	router := api.Router()
	router.Run("0.0.0.0:8080")
}
