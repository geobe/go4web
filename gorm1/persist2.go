package main

import (
	"github.com/jinzhu/gorm"
	//	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
	"github.com/geobe/go4j/poi"
	"github.com/geobe/go4web/gorm1/model"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	//	db, err := gorm.Open("sqlite3", "c:/usr/sqlitedata/gorm1.db")
	db, err := gorm.Open("postgres", "user=oosy dbname=gorm2 password=oosy2016 sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&model.City{}, &model.Destination{})

	var ct model.City
	var ds model.Destination

	for _, aCity := range poi.GermanCities {
		city := model.New(aCity)
		dest := model.Destination{Dest: city}
		// create dest sichert auch city in die DB
		// -> kaskadieren
		db.Create(&dest)
		//db.Create(&city)
	}

	db.First(&ct)
	fmt.Printf("City: %v\n", ct)

	db.First(&ds)
	fmt.Printf("Destination: %v\n\n", ds)

	var destinations []model.Destination
	// alle Destinations aus der Datenbank
	// in ein Array einlesen.
	// Das Feld Dest bleibt leer.
	db.Find(&destinations)
	fmt.Printf("len(destinations): %d\n"+
		"destinations[1]: %v\n"+
		"destinations[1].Dest.Name: %s\n",
		len(destinations), destinations[0], destinations[0].Dest.Name)

	// Preload holt die Objekte der
	// Assoziation "Dest" in das Objekt
	db.Preload("Dest").Find(&destinations)
	fmt.Printf("len(destinations): %d\n"+
		"destinations[1]: %v\n"+
		"destinations[1].Dest.Name: %s\n",
		len(destinations), destinations[0], destinations[0].Dest.Name)

	// Finde Destination, die auf "München"
	// zeigt, mit zwei Datenbankzugriffen
	var muc model.City
	var dmuc model.Destination
	db.First(&muc, "name = ?", "München")
	// First mit Query-Parameter aufrufen
	db.Preload("Dest").First(&dmuc, muc.DestinationID)
	fmt.Printf("dmuc.Dest.Name: %s\n", dmuc.Dest.Name)

	// Finde Destination, die auf "Berlin" zeigt, mit einem Datenbankzugriff
	// in JOIN muss die "fremde" Tabelle, zu der der Join läuft, vorn stehen
	var dmuc1 model.Destination
	db.Joins("JOIN cities On cities.destination_id = destinations.id"+
		" AND cities.name = ?", "Berlin").Preload("Dest").First(&dmuc1)
	fmt.Printf("dmuc.Dest.Name: %s\n", dmuc1.Dest.Name)

	db.Delete(model.City{})
	db.Delete(model.Destination{})

}
