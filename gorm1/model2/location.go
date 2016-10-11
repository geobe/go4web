// Package loc "verpackt" loc.Location für gorm.
// Dazu müssen alle Felder global sichtbar sein.
package model

import "github.com/geobe/go4j/loc"

type Location struct {
	Lat float64
	Lon float64
}

// Implementiere Interface Locator
func (loc Location) LatLon() (lat, lon float64) {
	lat = loc.Lat
	lon = loc.Lon
	return
}

// Berechne Abstand zwischen zwei Loc's,
// implementiert Interface Distancer.
func (l Location) Dist(lctr loc.Locator) float64 {
	return loc.Dist(l, lctr)
}
