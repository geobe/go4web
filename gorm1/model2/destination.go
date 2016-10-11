package model

type Destination struct {
	Model
	// Name ist sinnlos, aber ohne einen Wert
	// funktioniert Kaskadieren nicht
	Name   string
	CityID uint
}
