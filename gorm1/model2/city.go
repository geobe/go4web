package model2

import (
	"github.com/geobe/go4j/poi"
)

// model2.city "verpackt" poi.City für gorm.
// City macht alle Felder global sichtbar, sonst kann gorm
// damit nicht arbeiten.
type City struct {
	Model
	Location
	Name        string
	Inhabitants int
	Destination []Destination `gorm:"polymorphic:Dest;"`
}

// Konstruktor Function
func New(c poi.City) (city City) {
	lat, lon := c.LatLon()
	city = City{Location: Location{lat, lon}, Name: c.Name(),
		Inhabitants: c.Inhabitants()}
	return
}
