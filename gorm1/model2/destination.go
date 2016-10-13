package model2

// Reiseziel
type Destination struct {
	Model
	// beschreibt das Reiseziel für diesen Trip
	Reason   string
	TripId   uint
	DestID   uint
	DestType string
}
