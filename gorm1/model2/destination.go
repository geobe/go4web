package model2

// Reiseziel
type Destination struct {
	Model
	// beschreibt das Reiseziel für diesen Trip
	Reason   string
	Trip     Trip
	DestID   uint
	DestType string
}
