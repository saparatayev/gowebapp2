package models

import "errors"

var (
	ErrRequiredFirstname = errors.New("Firstame field is required")
	ErrRequiredLastname  = errors.New("Lastname field is required")
	ErrRequiredEmail     = errors.New("Email field is required")
	ErrRequiredPassword  = errors.New("Password field is required")
)

func IsEmpty(attr string) bool {
	if attr == "" {
		return true
	}

	return false
}

func ValidateNewUser(user User) (User, error) {
	if IsEmpty(user.Firstname) {
		return User{}, ErrRequiredFirstname
	}
	if IsEmpty(user.Lastname) {
		return User{}, ErrRequiredLastname
	}
	if IsEmpty(user.Email) {
		return User{}, ErrRequiredEmail
	}
	if IsEmpty(user.Password) {
		return User{}, ErrRequiredPassword
	}

	return user, nil
}
