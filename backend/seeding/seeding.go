package main

import (
	"fmt"
	"math/rand"

	api "github.com/AdamPekny/IIS/backend"
	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/utils"
	"github.com/brianvoe/gofakeit"
)

//go run ./backend/seeding/seeding.go pouzit raz na naseedovanie

func main() {
	gofakeit.Seed(69420)
	api.Migrate_all()
	db, _ := utils.Conn()
	// seed usertypes
	userTypes := []models.UserType{{CodeName: "admin"}, {CodeName: "technic"}, {CodeName: "driver"}}
	for _, v := range userTypes {
		db.Create(&v)
	}
	//seed users
	for i := 0; i < 10; i++ {
		user := models.User{
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Email:     gofakeit.Email(),
			BirthDate: gofakeit.Date(),
			Password:  gofakeit.Password(false, false, false, false, false, 5), // nieco ako hash lmao lmao
			Salt:      gofakeit.Password(false, false, false, false, false, 5),
			UserType:  userTypes[rand.Intn(3)],
		}
		db.Create(&user)
	}
	// seed vehicle types
	vehicleTypes := []models.VehicleType{{Type: "bus"}, {Type: "tram"}, {Type: "obrnena_dodavka"}}
	for _, v := range vehicleTypes {
		db.Create(&v)
	}
	//seed lines
	for i := 1; i <= 4; i++ {
		line := models.Line{
			Name:        fmt.Sprintf("Line %v", i),
			FinalStop:   gofakeit.StreetName(),
			InitialStop: gofakeit.StreetName(),
		}
		db.Create(&line)
	}
	lines := []models.Line{}
	db.Find(&lines) //TODO error
	// seed vehicles
	for i := 0; i < 10; i++ {
		vehicle := models.Vehicle{
			Capacity: uint(gofakeit.Number(5, 70)),
			// Brand
			// Image
			VehicleType: vehicleTypes[rand.Intn(3)],
		}
		db.Create(&vehicle)
	}
	vehicles := []models.Vehicle{}
	db.Find(&vehicles)
	// seed connections
	for i := 0; i < 10; i++ {
		connection := models.Connection{
			ArrivalTime:   gofakeit.Date(),
			DepartureTime: gofakeit.Date(),
			VehicleID:     vehicles[rand.Intn(len(vehicles)-1)].ID,
			LineName:      lines[rand.Intn(len(lines)-1)].Name,
		}
		db.Create(&connection)

	}
	stopsMap := make(map[string]bool)
	//seed stops
	for i := 0; i < len(lines); i++ {
		stop := models.Stop{
			Name: lines[i].InitialStop,
		}
		stopsMap[stop.Name] = true
		db.Create(&stop)
		stop = models.Stop{
			Name: lines[i].FinalStop,
		}
		stopsMap[stop.Name] = true
		db.Create(&stop)
	}
	for i := 0; i < 15; i++ {
		stop := models.Stop{
			Name: gofakeit.StreetName(),
		}
		db.Create(&stop)
	}
	//seed segments
	var stops []models.Stop

	db.Find(&stops) //TODO error

	for i := 0; i < len(stops)-1; i++ {
		if ok := stopsMap[stops[i].Name]; ok {
			continue
		} else {
			segment := models.Segment{
				Stop1: stops[i],
				Stop2: stops[i+1],
				Time:  uint(rand.Intn(3) + 2),
			}
			db.Create(&segment)
		}
	}
	segments := []models.Segment{}
	db.Find(&segments)
	for i := 0; i < len(lines); i++ {
		// initial segment
		segment := models.Segment{
			StopName1: lines[i].InitialStop,
			StopName2: segments[gofakeit.Number(0, 9)].StopName1,
			Time:      uint(rand.Intn(3) + 2),
		}
		db.Create(&segment)
		lines[i].Segments = append(lines[i].Segments, &segment)
		db.Save(&lines[i])
		// few other segments
		for j := 0; j < 2; j++ {
			findsegment := models.Segment{}
			db.Where(&models.Segment{StopName1: lines[i].Segments[j].StopName2}).First(&findsegment)
			lines[i].Segments = append(lines[i].Segments, &findsegment)
			db.Save(&lines[i])
		}
		// final segment
		segment = models.Segment{
			StopName1: segments[gofakeit.Number(10, 13)].StopName2,
			StopName2: lines[i].FinalStop,
			Time:      uint(rand.Intn(3) + 2),
		}
		db.Create(&segment)
		lines[i].Segments = append(lines[i].Segments, &segment)
		db.Save(&lines[i])
	}
}
