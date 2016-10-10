// Package city "verpackt" poi.City für gorm.
// Dazu müssen alle Felder global sichtbar sein.
package model

import (
	"github.com/geobe/go4j/poi"
)

// City macht alle Felder global sichtbar, sonst kann gorm
// damit nicht arbeiten.
type City struct {
	Model
	Location
	Name          string
	Inhabitants   int
	DestinationID uint
}

// Konstruktor Function
func New(c poi.City) (city City) {
	lat, lon := c.LatLon()
	city = City{Location: Location{lat, lon}, Name: c.Name(),
		Inhabitants: c.Inhabitants()}
	return
}
