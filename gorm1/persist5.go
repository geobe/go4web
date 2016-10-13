package main

import (
	"fmt"
	"github.com/geobe/go4j/poi"
	model "github.com/geobe/go4web/gorm1/model2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"strconv"
)

// Demoprogramm für polymorphe Assoziationen:
// für dieses Beispiel wird das geänderte model2 package verwendet
func main() {
	db, err := gorm.Open("postgres", "user=oosy dbname=gorm5 password=oosy2016 sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&model.City{}, &model.Attraction{}, &model.Destination{}, &model.Trip{}, &model.Person{})

	// Datenbank leeren
	db.Delete(model.Person{})
	db.Delete(model.Trip{})
	db.Delete(model.City{})
	db.Delete(model.Attraction{})
	db.Delete(model.Destination{})

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
		tx := db
		dest := model.Destination{Reason: "Karneval-" + strconv.Itoa(i)}
		c.Destination = append(c.Destination, dest)
		tx.Save(&c)
	}

	dest := model.Destination{Reason: "skurriles Schloß"}
	var att model.Attraction
	db.First(&att, "name like 'Neuschw%'")
	att.Destination = append(att.Destination, dest)
	db.Save(&att)

	db.Find(&dests)
	for _, dest := range dests {
		var city model.City
		var attr model.Attraction
		fmt.Printf("Reiseziel %s: ", dest.Reason)
		// ausführliche Variante:
		// Polymorphes Objekt vollständig lesen
		db.First(&attr, dest.DestID)
		db.First(&city, dest.DestID)
		if city.ID != 0 {
			fmt.Printf("City %s\n", city.Name)
		} else {
			fmt.Printf("Attraction %s\n", attr.Name)
		}
		// kompakte Variante: nur die Werte
		// lesen, die gebraucht werden
		var any struct{ Name string }
		db.Table(dest.DestType).Where("ID = ?", dest.DestID).Scan(&any)
		fmt.Printf("\t%s\n", any.Name)
	}

	kirk.Trips[0].Destinations = dests

	db.Save(&kirk)

	// query
	var kirki model.Person

	db.Preload("Trips").
		Preload("Trips.Destinations").
		First(&kirki, kirk.ID)

	fmt.Printf("Person %s, %d Trips, 1. Trip %s hat %d Stationen:\n",
		kirki.Name, len(kirki.Trips), kirki.Trips[0].Comment,
		len(kirki.Trips[0].Destinations))
	for _, kdest := range kirki.Trips[0].Destinations {
		var any struct {
			Description string
			Name        string
		}
		db.Table(kdest.DestType).Where("ID = ?", kdest.DestID).Scan(&any)
		fmt.Printf("\t%s: %s %s\n", kdest.Reason,
			any.Description, any.Name)
	}

	//fmt.Println(kirk)
	//fmt.Println(kiki)

}
