package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SpecieID primitive.ObjectID

type Specie struct {
	ID                  primitive.ObjectID        `bson:"_id" json:"id,omitempty" example:"1"`
	Name                string                    `json:"name" validate:"required"`
	BaseStatus          Status                    `json:"baseStatus" validate:"required"`
	LearnableTechniques []TechniqueAndRequirement `json:"learnableTechniques" validate:"required"`
}
