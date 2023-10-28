package main

import (
	"fmt"
	"math/rand"

	api "github.com/AdamPekny/IIS/backend"
	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/utils"
	"github.com/brianvoe/gofakeit/v6"
)

// go run ./backend/seeding/seeding.go pouzit raz na naseedovanie
func init() {
	utils.LoadEnvVariables()
	utils.Conn()
}

func main() {
	gofakeit.Seed(69420)
	api.Migrate_all()
	//users need to be seeded manually
	//seed vehicle types
	vehicleTypes := []models.VehicleType{{Type: "bus"}, {Type: "tram"}, {Type: "obrnena_dodavka"}}
	for _, v := range vehicleTypes {
		utils.DB.Create(&v)
	}
	//seed vehicles
	for i := 0; i < 10; i++ {
		vehicle := models.Vehicle{
			Capacity:     uint(gofakeit.Number(5, 70)),
			Registration: gofakeit.LetterN(3) + string(gofakeit.Number(1000, 9999)),
			// Brand
			// Image
			VehicleType: vehicleTypes[rand.Intn(3)],
		}
		utils.DB.Create(&vehicle)
	}
	//seed stops
	for i := 0; i < 22; i++ {
		stop := models.Stop{
			Name: gofakeit.StreetName(),
		}
		utils.DB.Create(&stop)
	}
	var stops []models.Stop
	utils.DB.Find(&stops)
	//seed lines
	for i := 1; i <= 2; i++ {
		line := models.Line{
			Name:        fmt.Sprintf("Line %v", i),
			FinalStop:   stops[gofakeit.Number(0, 9)].Name,
			InitialStop: stops[gofakeit.Number(10, 19)].Name,
		}
		utils.DB.Create(&line)
	}
	lines := []models.Line{}
	utils.DB.Find(&lines)

	vehicles := []models.Vehicle{}
	utils.DB.Find(&vehicles)
	// seed segments
	for i := 0; i < len(lines); i++ {
		initial_segment := models.Segment{
			StopName1: lines[i].InitialStop,
			StopName2: stops[gofakeit.Number(0, len(stops)-1)].Name,
			Time:      uint(rand.Intn(3) + 1),
		}
		utils.DB.Create(&initial_segment)
		lines[i].Segments = append(lines[i].Segments, &initial_segment)
		utils.DB.Save(&lines[i])
		for j := 0; j < 3; j++ {
			next_segment := models.Segment{
				StopName1: lines[i].Segments[len(lines[i].Segments)-1].StopName2,
				StopName2: stops[gofakeit.Number(0, len(stops)-1)].Name,
				Time:      uint(rand.Intn(3) + 1),
			}
			utils.DB.Create(&next_segment)
			lines[i].Segments = append(lines[i].Segments, &next_segment)
			utils.DB.Save(&lines[i])
		}
		final_segment := models.Segment{
			StopName1: lines[i].Segments[len(lines[i].Segments)-1].StopName2,
			StopName2: lines[i].FinalStop,
			Time:      uint(rand.Intn(3) + 1),
		}
		utils.DB.Create(&final_segment)
		lines[i].Segments = append(lines[i].Segments, &final_segment)
		utils.DB.Save(&lines[i])
	}
	for i := 0; i < 8; i++ {
		connection := models.Connection{
			DepartureTime: gofakeit.Date(),
			VehicleID:     vehicles[rand.Intn(len(vehicles)-1)].ID,
			LineName:      lines[i%2].Name,
			Dirrection:    i%2 == 1,
		}
		utils.DB.Create(&connection)

	}
}
