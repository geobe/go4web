package model

import (
	"github.com/geobe/go4j/poi"
)

// model.City "verpackt" poi.City für gorm.
// model.City macht alle Felder global sichtbar,
// sonst kann gorm damit nicht arbeiten.
type City struct {
	Model
	Location
	Name          string
	Inhabitants   int
	DestinationID uint
	Trips         []Trip `gorm:"many2many:trip_city;"`
}

// Konstruktor Function
func New(c poi.City) (city City) {
	lat, lon := c.LatLon()
	city = City{Location: Location{lat, lon}, Name: c.Name(),
		Inhabitants: c.Inhabitants()}
	return
}

/*
func (c City) BeforeDelete() {
	if c.ID == 0 {
		fmt.Printf("Deleting all cities\n")
	} else {
		fmt.Printf("Deleting City %s\n", c.Name)
	}
}
*/
