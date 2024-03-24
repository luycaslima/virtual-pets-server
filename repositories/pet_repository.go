package repositories

import (
	"context"

	"github.com/luycaslima/virtual-pets-server/models"
)

type PetRepository interface {
	InsertANewPet(ctx context.Context, in *models.Pet) error
	UpdatePetData(ctx context.Context, updatedPet *models.Pet) error
	GetPetPublicData(ctx context.Context, petId string) (*models.PetPublicData, error)
	GetPetAllData(ctx context.Context, petId string) (*models.Pet, string, int, error)
}
