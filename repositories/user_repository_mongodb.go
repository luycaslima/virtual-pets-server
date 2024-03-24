package repositories

import (
	"context"

	"github.com/luycaslima/virtual-pets-server/database"
	"github.com/luycaslima/virtual-pets-server/dto"
	"github.com/luycaslima/virtual-pets-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userMongoDbRepository struct {
	collection *mongo.Collection
}

func NewUserMongoDbRepository(db *mongo.Client) userMongoDbRepository {
	collection := database.GetCollectionFromDB(db, "users")
	return userMongoDbRepository{collection: collection}
}

func (r *userMongoDbRepository) InsertANewUser(ctx context.Context, in *models.User) error {
	_, err := r.collection.InsertOne(ctx, in)
	if err != nil {
		return err
	}
	return nil
}

func (r *userMongoDbRepository) CheckUser(ctx context.Context, in *dto.UserLoginRequest) (*models.User, error) {
	var foundUser models.User
	//Search user by username
	if err := r.collection.FindOne(ctx, bson.M{"publicdata.username": in.Username}).Decode(&foundUser); err != nil {
		return nil, err
	}

	return &foundUser, nil
}

func (r *userMongoDbRepository) FindAnUserById(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	var foundedUser models.User

	if err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&foundedUser); err != nil {
		return nil, err
	}
	return &foundedUser, nil
}

func (r *userMongoDbRepository) FindAnUserByUsername(ctx context.Context, username string) (*models.UserPublicData, error) {
	var foundedUser models.User

	if err := r.collection.FindOne(ctx, bson.M{"publicdata.username": username}).Decode(&foundedUser); err != nil {
		return nil, err
	}
	return &foundedUser.PublicData, nil
}

func (r *userMongoDbRepository) LinkAPetToAnUser(ctx context.Context, petId primitive.ObjectID, user *models.User) error {

	_, err := r.collection.UpdateByID(ctx, user.ID, bson.M{"$push": bson.M{"publicdata.pets": petId}})
	if err != nil {
		return err
	}

	return nil
}
