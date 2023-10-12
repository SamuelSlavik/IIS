package main

import (
	api "github.com/AdamPekny/IIS/backend"
	"github.com/AdamPekny/IIS/backend/utils"
)

func main() {
	utils.Migrate_all()
	api.Router()
}
