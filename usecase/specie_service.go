package usecase

import (
	"context"

	"github.com/luycaslima/virtual-pets-server/dto"
	"github.com/luycaslima/virtual-pets-server/models"
	"github.com/luycaslima/virtual-pets-server/repositories"
)

type SpecieService struct {
	specieRepository repositories.SpecieRepository
}

func NewSpecieService(repository repositories.SpecieRepository) *SpecieService {
	return &SpecieService{
		specieRepository: repository,
	}
}

func (s *SpecieService) RegisterANewSpecie(ctx context.Context, in *models.Specie) error {
	return s.specieRepository.InsertANewSpecie(ctx, in)
}

func (s *SpecieService) GetSpecieDetails(ctx context.Context, specieId string) (*models.Specie, string, int, error) {
	specie, msg, statusCode, err := s.specieRepository.GetASpecieDetails(ctx, specieId)
	if err != nil {
		return nil, msg, statusCode, err
	}
	return specie, msg, statusCode, nil
}

func (s *SpecieService) GetListOfSpecies(ctx context.Context) (*dto.ListSpeciesResponse, error) {
	listOfSpecies, err := s.specieRepository.GetListOfOverviewSpecies(ctx)
	if err != nil {
		return nil, err
	}
	return listOfSpecies, nil
}
