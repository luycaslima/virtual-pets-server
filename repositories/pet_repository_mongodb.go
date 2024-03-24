package repositories

import (
	"context"
	"errors"
	"net/http"

	"github.com/luycaslima/virtual-pets-server/database"
	"github.com/luycaslima/virtual-pets-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type petMongoDbRepository struct {
	collection *mongo.Collection
}

func NewPetMongoDbRepository(db *mongo.Client) petMongoDbRepository {
	collection := database.GetCollectionFromDB(db, "pets")
	return petMongoDbRepository{collection}
}

func (r *petMongoDbRepository) InsertANewPet(ctx context.Context, in *models.Pet) error {
	_, err := r.collection.InsertOne(ctx, in)
	if err != nil {
		return err
	}
	return nil
}

// TODO implement
func (r *petMongoDbRepository) UpdatePetData(ctx context.Context, updatedPet *models.Pet) error {
	return errors.New("not implemented")
}

// TODO implement
func (r *petMongoDbRepository) GetPetPublicData(ctx context.Context, petId string) (*models.PetPublicData, error) {
	return nil, errors.New("not implemented")
}

func (r *petMongoDbRepository) GetPetAllData(ctx context.Context, petId string) (*models.Pet, string, int, error) {
	var foundedPet models.Pet

	id, err := primitive.ObjectIDFromHex(petId)
	if err != nil {
		return nil, "Error converting id", http.StatusInternalServerError, err
	}

	if err = r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&foundedPet); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, "Pet Not Founded", http.StatusNotFound, err
		}
	}

	return &foundedPet, "Founded!", http.StatusOK, nil
}
