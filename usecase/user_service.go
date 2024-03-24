package usecase

import (
	"context"
	"net/http"
	"time"

	"github.com/luycaslima/virtual-pets-server/auth"
	"github.com/luycaslima/virtual-pets-server/dto"
	"github.com/luycaslima/virtual-pets-server/models"
	"github.com/luycaslima/virtual-pets-server/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TODO not interfaced implemented, thefore use pointer

type UserService struct {
	userRepository   repositories.UserRepository
	specieRepository repositories.SpecieRepository
	petRepository    repositories.PetRepository
}

func NewUserService(userRep repositories.UserRepository, specieRep repositories.SpecieRepository, petRep repositories.PetRepository) *UserService {
	return &UserService{
		userRepository:   userRep,
		specieRepository: specieRep,
		petRepository:    petRep,
	}
}

func (s *UserService) RegisterAnUser(ctx context.Context, in *dto.UserRegistrationRequest) error {
	//Encrypt the password
	password, passwordErr := auth.HashPassword(string(in.Password))
	if passwordErr != nil {
		return passwordErr
	}

	newUser := models.User{
		ID: primitive.NewObjectID(),
		PublicData: models.UserPublicData{
			Username:    in.Username,
			Pets:        []primitive.ObjectID{},
			Vivariums:   []primitive.ObjectID{},
			CreatedDate: time.Now().Format("2006-01-02 15:04:05"),
		},
		PrivateData: models.UserPrivateData{
			Email: in.Email,
			Money: 1000,
		},
		InternalData: models.UserInternalData{
			Password: password,
		},
	}

	return s.userRepository.InsertANewUser(ctx, &newUser)
}

func (s *UserService) LoginAnUser(ctx context.Context, in *dto.UserLoginRequest) (*http.Cookie, string, int, error) {
	user, err := s.userRepository.CheckUser(ctx, in)
	if err != nil {
		return nil, "User not Found", http.StatusNotFound, err
	}

	if isPasswordCorrect := auth.CheckPassword(in.Password, user.InternalData.Password); !isPasswordCorrect {
		return nil, "Wrong Password", http.StatusUnauthorized, err
	}

	//Generate and return a token
	token, expirationDate, err := auth.CreateJWTToken(user.ID.Hex())
	if err != nil {
		return nil, "Error creating jwt token", http.StatusInternalServerError, err
	}

	cookie := http.Cookie{
		SameSite: http.SameSiteNoneMode,
		Name:     "jwt",
		Value:    token,
		Expires:  expirationDate,
		HttpOnly: true,
	}

	return &cookie, "Logged successfully", http.StatusOK, nil
}

func (s *UserService) GetUserPublicDataByUsername(ctx context.Context, username string) (*models.UserPublicData, error) {
	user, err := s.userRepository.FindAnUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) CreateAPet(ctx context.Context, userId string, in *dto.CreatePetRequest) (*models.Pet, string, int, error) {
	userID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, "Error on id conversion", http.StatusInternalServerError, err
	}

	//Check if the user exists
	foundedUser, userErr := s.userRepository.FindAnUserById(ctx, userID)
	if userErr != nil {
		return nil, "Error Finding User", http.StatusNotFound, err
	}

	foundedSpecie, msg, statusCode, specieErr := s.specieRepository.GetASpecieDetails(ctx, in.SpecieID)
	if specieErr != nil {
		return nil, msg, statusCode, specieErr
	}
	//TODO integrate techniques
	newPet := models.Pet{
		ID: primitive.NewObjectID(),
		PublicData: models.PetPublicData{
			PetName:    in.PetName,
			Birthday:   time.Now().String(),
			Techniques: []string{},
			Status:     foundedSpecie.BaseStatus,
			Specie: models.SpecieOverview{
				ID:   foundedSpecie.ID,
				Name: foundedSpecie.Name,
			},
		},
		InternalData: models.PetInternalData{
			OwnerID:   foundedUser.ID,
			Hunger:    0,
			Happiness: 100,
			Cleanness: 100,
		},
	}

	err = s.petRepository.InsertANewPet(ctx, &newPet)
	if err != nil {
		return nil, "Error inserting a new Pet", http.StatusInternalServerError, err
	}

	err = s.userRepository.LinkAPetToAnUser(ctx, newPet.ID, foundedUser)

	if err != nil {
		return nil, "Error linking a Pet to a User", http.StatusInternalServerError, err
	}

	return &newPet, "Created successfully", http.StatusOK, nil
}
