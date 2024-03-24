package repositories

import (
	"context"

	"github.com/luycaslima/virtual-pets-server/dto"
	"github.com/luycaslima/virtual-pets-server/models"
)

type SpecieRepository interface {
	InsertANewSpecie(ctx context.Context, in *models.Specie) error
	GetListOfOverviewSpecies(ctx context.Context) (*dto.ListSpeciesResponse, error) //TODO remove and only use by filter
	GetListOfSpeciesByFilter(ctx context.Context, query string) (*dto.ListSpeciesResponse, error)
	GetASpecieDetails(ctx context.Context, specieId string) (*models.Specie, string, int, error)
}
