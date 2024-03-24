package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//TODO have an enum for separate each evolutive phase of the specie

// TODO need to have
// SpriteURL
// DetailSpriteURL

type Specie struct {
	ID                  primitive.ObjectID        `bson:"_id" json:"id" example:"1"`
	Name                string                    `json:"name" validate:"required"`
	Description         string                    `json:"description" validate:"required"`
	BaseStatus          Status                    `json:"baseStatus" validate:"required"`
	LearnableTechniques []TechniqueAndRequirement `json:"learnableTechniques" validate:"required"`
}

// Have an sprite img
type SpecieOverview struct {
	ID   primitive.ObjectID `bson:"_id" json:"id" example:"1"`
	Name string             `json:"specieName" validate:"required"`
}
