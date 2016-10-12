package model2

type Person struct {
	Model
	Name  string
	Trips []Trip
}

var SomePersons = []Person{
	{Name: "Captain Kirk"},
	{Name: "Lara Croft"},
}
