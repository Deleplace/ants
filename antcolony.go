package ants

// An Ant is a struct with a few arbitrary fields.
type Ant struct {
	Field1 int
	Field2 string
	Field3 int
	Field4 string
	Field5 int
	Field6 string
	Field7 int
	Field8 string
}

// An AntColony is a slice of Ants.
type AntColony []Ant

// A DataOrientedAntColony is an alternative representation of
// an AntColony, where all the values of each fields are stored
// in a contiguous slice, in a data-oriented manner.
type DataOrientedAntColony struct {
	Field1 []int
	Field2 []string
	Field3 []int
	Field4 []string
	Field5 []int
	Field6 []string
	Field7 []int
	Field8 []string
}
