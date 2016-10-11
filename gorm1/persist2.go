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

	for _, aCity := range poi.GermanCities {
		city := model.New(aCity)
		dest := model.Destination{Dest: city}
		//fmt.Printf("City: %v\n", city)
		// create dest sichert auch city in die DB
		// -> kaskadieren
		db.Create(&dest)
		//db.Create(&c)
	}

	db.First(&ct)
	fmt.Printf("City: %v\n", ct)

	var destinations []model.Destination

	// Dest bleibt leer
	db.Find(&destinations)
	fmt.Printf("len(destinations): %d\n"+
		"destinations[1]: %v\n"+
		"destinations[1].Dest.Name: %s\n",
		len(destinations), destinations[0], destinations[0].Dest.Name)

	// Preload holt die Objekte der Relationship "Dest" in das Objekt
	db.Preload("Dest").Find(&destinations)
	fmt.Printf("len(destinations): %d\n"+
		"destinations[1]: %v\n"+
		"destinations[1].Dest.Name: %s\n",
		len(destinations), destinations[0], destinations[0].Dest.Name)

	// Finde Destination, die auf "München" zeigt
	var muc model.City
	var dmuc model.Destination
	db.First(&muc, "name = ?", "München")
	db.Preload("Dest").First(&dmuc, muc.ID)
	fmt.Printf("dmuc.Dest.Name: %s\n", dmuc.Dest.Name)

	db.Delete(model.City{})
	db.Delete(model.Destination{})

}
