package repositories

import (
	"context"

	"github.com/luycaslima/virtual-pets-server/dto"
	"github.com/luycaslima/virtual-pets-server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	CheckUser(ctx context.Context, in *dto.UserLoginRequest) (*models.User, error)
	InsertANewUser(ctx context.Context, in *models.User) error
	FindAnUserByUsername(ctx context.Context, username string) (*models.UserPublicData, error)
	LinkAPetToAnUser(ctx context.Context, petId primitive.ObjectID, user *models.User) error
	FindAnUserById(ctx context.Context, id primitive.ObjectID) (*models.User, error)
}
