package models

type TechniqueRange uint

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
	Technique    string `json:"techniqueName" validate:"required"` //TODO change to an id from the technique
	Requirements Status `json:"requirements" validate:"required"`
}
