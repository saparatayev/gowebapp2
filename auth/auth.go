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
	err := validateFields(email, password)

	if err != nil {
		return models.User{}, err
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

func validateFields(email, password string) error {
	if models.IsEmpty(email) || models.IsEmpty(password) {
		return ErrEmptyFields
	}

	if !models.IsEmail(email) {
		return models.ErrInvalidEmail
	}

	return nil
}
