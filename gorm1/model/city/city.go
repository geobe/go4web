// Package city "verpackt" poi.City für gorm.
// Dazu müssen alle Felder global sichtbar sein.
package city

import (
	"github.com/geobe/go4j/poi"
	"github.com/geobe/go4web/gorm1/model/loc"
	"github.com/jinzhu/gorm"
)

// City macht alle Felder global sichtbar, sonst kann gorm
// damit nicht arbeiten.
type City struct {
	gorm.Model
	loc.Loc
	Name        string
	Inhabitants int
}

// Konstruktor Function
func New(c poi.City) (city City) {
	lat, lon := c.LatLon()
	city = City{Loc: loc.Loc{lat, lon}, Name: c.Name(),
		Inhabitants: c.Inhabitants()}
	return
}
