package repositories

import (
	"context"
	"errors"
	"net/http"

	"github.com/luycaslima/virtual-pets-server/database"
	"github.com/luycaslima/virtual-pets-server/dto"
	"github.com/luycaslima/virtual-pets-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type specieMongoDbRepository struct {
	collection *mongo.Collection
}

func NewSpecieMongoDbRepository(db *mongo.Client) specieMongoDbRepository {
	collection := database.GetCollectionFromDB(db, "species")
	return specieMongoDbRepository{collection: collection}
}

func (r *specieMongoDbRepository) InsertANewSpecie(ctx context.Context, in *models.Specie) error {
	_, err := r.collection.InsertOne(ctx, in)
	if err != nil {
		return err
	}
	return nil
}

func (r *specieMongoDbRepository) GetListOfOverviewSpecies(ctx context.Context) (*dto.ListSpeciesResponse, error) {

	var results dto.ListSpeciesResponse

	//find records
	//pass these options to the Find method
	//findOptions := options.Find()
	//Set the limit of the number of record to find
	//findOptions.SetLimit(5)

	//which fields i want to retrive
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "name", Value: 1},
	}

	cursor, err := r.collection.Find(ctx, bson.M{}, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var specieOverview models.SpecieOverview
		if err := cursor.Decode(&specieOverview); err != nil {
			return nil, err
		}
		results.Species = append(results.Species, &specieOverview)
	}
	return &results, nil
}

func (r *specieMongoDbRepository) GetListOfSpeciesByFilter(ctx context.Context, query string) (*dto.ListSpeciesResponse, error) {
	return nil, errors.New("not implemented")
}

func (r *specieMongoDbRepository) GetASpecieDetails(ctx context.Context, specieId string) (*models.Specie, string, int, error) {
	var foundedSpecie models.Specie
	id, err := primitive.ObjectIDFromHex(specieId)

	if err != nil {
		return nil, "Failed Conversion on the Server", http.StatusInternalServerError, err
	}

	if err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&foundedSpecie); err != nil {
		return nil, "Species not found!", http.StatusNotFound, err
	}

	return &foundedSpecie, "Founded!", http.StatusOK, nil
}
