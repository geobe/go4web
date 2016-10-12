package model

import "time"

type Trip struct {
	Model
	Comment  string
	Start    time.Time
	PersonID uint
	Cities   []City `gorm:"many2many:trip_city;"`
}

var SomeTrips = []Trip{
	{Comment: "Karneval", Start: tx("11.11.2016")},
	{Comment: "Neujahrstrip", Start: tx("1.1.2017")},
	{Comment: "Sommerreise", Start: tx("12.7.2017")},
}

func tx(timestring string) (t time.Time) {
	t, _ = time.Parse("2.1.2006", timestring)
	return
}
