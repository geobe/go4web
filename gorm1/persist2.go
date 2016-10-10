package main

import (
	"github.com/jinzhu/gorm"
	//	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
	"github.com/geobe/go4j/poi"
	"github.com/geobe/go4web/gorm1/model/city"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	//	db, err := gorm.Open("sqlite3", "c:/usr/sqlitedata/gorm1.db")
	db, err := gorm.Open("postgres", "user=oosy dbname=gorm2 password=oosy2016 sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&city.City{})

	var ct city.City

	for _, aCity := range poi.GermanCities {
		c := city.New(aCity)
		fmt.Printf("City: %v\n", c)
		db.Create(&c)
	}

	db.First(&ct)

	fmt.Printf("City: %v\n", ct)

}
