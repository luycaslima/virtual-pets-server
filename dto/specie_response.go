package dto

import "github.com/luycaslima/virtual-pets-server/models"

type ListSpeciesResponse struct {
	Species []*models.SpecieOverview
}
