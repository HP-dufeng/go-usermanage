package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dufeng/usermanager/common"
	"github.com/dufeng/usermanager/controllers/dtos"
	"github.com/dufeng/usermanager/data"
	"github.com/dufeng/usermanager/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var inputData dtos.RegisterInput
	// Decode the incoming params json
	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid User data",
			500,
		)

		return
	}

	context := NewContext()
	defer context.Close()

	c := context.DbCollection("users")
	repo := &data.UserRepository{c}

	// Insert user
	userEntity := inputData.MapToUserEntity()
	repo.CreateUser(&userEntity)

	// Write response
	j, err := json.Marshal(dtos.MapToRegisterOutput(&userEntity))
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var inputData dtos.LoginInput
	var token string

	// Decode the incoming params json
	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Login data",
			500,
		)
		return
	}

	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{c}

	// Authenticate the login user
	user, err := repo.Login(models.User{
		Email:    inputData.Email,
		Password: inputData.Password,
	})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid login credentials",
			401,
		)
		return
	}

	//if login is successful
	// Generate JWT token
	token, err = common.GenerateJWT(user.Email, "member")
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Eror while generating the access token",
			500,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	j, err := json.Marshal(dtos.LoginOutput{Token: token})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}
