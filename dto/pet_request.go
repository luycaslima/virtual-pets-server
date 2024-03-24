package dto

type CreatePetRequest struct {
	PetName  string `json:"petName" validate:"required"`
	SpecieID string `json:"specieID" validate:"required"`
}
