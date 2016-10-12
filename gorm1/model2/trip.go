package model2

import "time"

type Trip struct {
	Model
	Comment      string
	Start        time.Time
	PersonID     uint
	Destinations []Destination
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
