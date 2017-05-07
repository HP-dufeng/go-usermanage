package dtos

import (
	"github.com/dufeng/usermanager/models"
	"gopkg.in/mgo.v2/bson"
)

type (

	// RegisterInput /user/register
	RegisterInput struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	// RegisterOutput /user/register
	RegisterOutput struct {
		ID        bson.ObjectId `json:"id"`
		FirstName string        `json:"firstname"`
		LastName  string        `json:"lastname"`
		Email     string        `json:"email"`
	}

	// LoginInput /user/login
	LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// LoginOutput response for /user/login
	LoginOutput struct {
		Token string `json:"token"`
	}
)

// MapToUserEntity convert RegisterInput to User entity
func (m *RegisterInput) MapToUserEntity() models.User {
	return models.User{
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Email:     m.Email,
		Password:  m.Password,
	}
}

// MapToRegisterOutput convert User entity to RegisterOutput
func MapToRegisterOutput(m *models.User) RegisterOutput {
	return RegisterOutput{
		ID:        m.Id,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Email:     m.Email,
	}
}
