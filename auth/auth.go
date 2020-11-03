package auth

import (
	"errors"
	"gowebapp2/models"
	"gowebapp2/utils"
)

var (
	ErrEmailNotFound   = errors.New("No such email")
	ErrInvalidPassword = errors.New("Password is not valid")
	ErrEmptyFields     = errors.New("Empty fields")
)

func Signin(email, password string) (models.User, error) {
	if models.IsEmpty(email) || models.IsEmpty(password) {
		return models.User{}, ErrEmptyFields
	}

	user, err := models.GetUserByEmail(email)
	if err != nil {
		return user, err
	}

	if user.Id == 0 {
		return user, ErrEmailNotFound
	}

	err = utils.VerifyPassword(user.Password, password)
	if err != nil {
		return models.User{}, ErrInvalidPassword
	}

	return user, nil
}
