package models

type TechniqueRange int64

const (
	ShortDistance TechniqueRange = iota
	MediumDistance
	LongDistance
)

// TODO use this in future
//
//	type Technique struct {
//		techID primitive.ObjectID
//		Name   string
//
// Cost int
// Hit
// DMG
// Atackrange
// }

type TechniqueAndRequirement struct {
	Technique    string //TODO change to an id from the technique
	Requirements []Atributte
}
