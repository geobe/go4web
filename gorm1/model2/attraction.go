package model2

import (
	"github.com/jinzhu/gorm"
)

type Attraction struct {
	gorm.Model
	Location
	Name        string
	Description string
	Stars       int
	Destination []Destination `gorm:"polymorphic:Dest;"`
}

var GermanAttractions = []Attraction{
	{Location: Location{7.884062, 54.182850}, Name: "Helgoland", Description: "rote Nordseeinsel", Stars: 3},
	{Location: Location{10.750085, 47.557537}, Name: "Neuschwanstein", Description: "Märchenschloß eines verrückten Königs", Stars: 2},
	{Location: Location{13.464394, 51.156154}, Name: "Meißen", Description: "Erste deutsche Porzellanmanufaktur", Stars: 4},
}
