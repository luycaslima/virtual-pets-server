package dto

type UserRegistrationRequest struct {
	Username string `json:"username" validate:"required" example:"ronald123"`
	Password string `json:"password" validate:"required" example:"ajksdnlaksjdmasnd"`
	Email    string `json:"email" validate:"required" example:"test@mail.com"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required" example:"ronald123"`
	Password string `json:"password" validate:"required" example:"ajksdnlaksjdmasnd"`
}

type UserPetLinkRequest struct {
	UserId string `json:"userId"`
	PetId  string `json:"petId"`
}
