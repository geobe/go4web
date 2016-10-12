package main

import (
	"fmt"
	"github.com/geobe/go4web/gorm1/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "user=oosy dbname=gorm3 password=oosy2016 sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&model.Trip{}, &model.Person{})

	kirk := model.SomePersons[0]
	lara := model.SomePersons[1]
	kirk.Trips = append(kirk.Trips, model.SomeTrips...)
	lara.Trips = append(lara.Trips, model.SomeTrips[1])

	db.Save(&kirk)
	db.Save(&lara)

	// query
	var kiki, lada model.Person

	db.First(&kiki, kirk.ID)
	db.Preload("Trips").First(&lada, lara.ID)

	fmt.Println(kirk)
	fmt.Println(kiki)

	fmt.Println(lara)
	fmt.Println(lada)

	db.Delete(model.Person{})
	db.Delete(model.Trip{})

}
