package main

import (
	"fmt"
	"github.com/geobe/go4j/poi"
	"github.com/geobe/go4web/gorm1/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "user=oosy dbname=gorm4 password=oosy2016 sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&model.City{}, &model.Destination{}, &model.Trip{}, model.Person{})

	for _, aCity := range poi.GermanCities {
		city := model.New(aCity)
		dest := model.Destination{Dest: city}
		// create dest sichert auch city in die DB
		// -> kaskadieren
		db.Create(&dest)
	}

	kirk := model.SomePersons[0]
	kirk.Trips = append(kirk.Trips, model.SomeTrips[0], model.SomeTrips[2])

	var dests []model.City
	db.Find(&dests, "name in ('München', 'Köln', 'Düsseldorf')")
	kirk.Trips[0].Cities = append(kirk.Trips[0].Cities, dests...)
	db.Find(&dests, "name in ('Zwickau', 'Leipzig', 'Dresden')")
	kirk.Trips[1].Cities = append(kirk.Trips[1].Cities, dests...)

	db.Save(&kirk)

	// query
	var kirki model.Person

	db.Preload("Trips").First(&kirki, kirk.ID)
	db.Preload("Cities").First(&kirki.Trips[0])

	fmt.Printf("Person %s, %d Trips, 1. Trip %s hat %d Stationen: ",
		kirki.Name, len(kirki.Trips), kirki.Trips[0].Comment,
		len(kirki.Trips[0].Cities))
	for _, city := range kirki.Trips[0].Cities {
		fmt.Printf("%s, ", city.Name)
	}
	println()

	//fmt.Println(kirk)
	//fmt.Println(kiki)

	db.Delete(model.Person{})
	db.Delete(model.Trip{})
	db.Delete(model.City{})
	db.Delete(model.Destination{})
	db.Exec("delete from trips_cities")
}
