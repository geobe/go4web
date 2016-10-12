package main

import (
	"fmt"
	"github.com/geobe/go4j/poi"
	model "github.com/geobe/go4web/gorm1/model2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"strconv"
)

// für dieses Beispiel wird das geänderte model2 package verwendet
func main() {
	db, err := gorm.Open("postgres", "user=oosy dbname=gorm5 password=oosy2016 sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&model.City{}, &model.Attraction{}, &model.Destination{}, &model.Trip{}, &model.Person{})

	for _, aCity := range poi.GermanCities {
		city := model.New(aCity)
		db.Create(&city)
	}

	for _, attr := range model.GermanAttractions {
		db.Create(&attr)
	}

	kirk := model.SomePersons[0]
	kirk.Trips = append(kirk.Trips, model.SomeTrips[0], model.SomeTrips[2])

	var dests []model.Destination
	var cities []model.City
	db.Find(&cities, "name in ('Köln', 'München', 'Düsseldorf')")
	for i, c := range cities {
		tx := db.Begin()
		dest := model.Destination{Reason: "Karneval " + strconv.Itoa(i)}
		tx.Save(&dest)
		c.Destination = append(c.Destination, dest)
		tx.Save(&c)
		dests = append(dests, dest)
		tx.Commit()
	}

	fmt.Printf("dests %v\n", dests)
	fmt.Printf("cities %v\n", cities)

	kirk.Trips[0].Destinations = append(kirk.Trips[0].Destinations, dests...)
	//db.Joins("JOIN cities On cities.id = destinations.city_id"+
	//	" AND cities.name in (?)", []string{"Zwickau", "Leipzig", "Dresden", "Bremen"}).Find(&dests)
	//kirk.Trips[1].Destinations = append(kirk.Trips[1].Destinations, dests...)

	db.Save(&kirk)

	// query
	var kirki model.Person

	db.Preload("Trips").First(&kirki, kirk.ID)
	db.Preload("Destinations").Find(&kirki.Trips)

	fmt.Printf("Person %s, %d Trips, 2. Trip %s hat %d Stationen: ",
		kirki.Name, len(kirki.Trips), kirki.Trips[1].Comment,
		len(kirki.Trips[1].Destinations))
	for _, aDest := range kirki.Trips[1].Destinations {
		var city model.City
		db.Find(&city, aDest.DestID)
		fmt.Printf(" %s,", city.Name)
	}
	println()

	//fmt.Println(kirk)
	//fmt.Println(kiki)

	db.Delete(model.Person{})
	db.Delete(model.Trip{})
	db.Delete(model.City{})
	db.Delete(model.Destination{})
	db.Exec("delete from trips_destinations")
}
