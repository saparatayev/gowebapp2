package models

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

var (
	ErrRequiredFirstname = errors.New("Firstame field is required")
	ErrRequiredLastname  = errors.New("Lastname field is required")
	ErrRequiredEmail     = errors.New("Email field is required")
	ErrInvalidEmail      = errors.New("Email is invalid")
	ErrRequiredPassword  = errors.New("Password field is required")
)

func IsEmpty(attr string) bool {
	if attr == "" {
		return true
	}

	return false
}

func Trim(attr string) string {
	return strings.TrimSpace(attr)
}

func IsEmail(email string) bool {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return false
	}

	return true
}

func ValidateNewUser(user User) (User, error) {
	user.Firstname = Trim(user.Firstname)
	user.Lastname = Trim(user.Lastname)
	user.Email = Trim(strings.ToLower(user.Email))

	if IsEmpty(user.Firstname) {
		return User{}, ErrRequiredFirstname
	}
	if IsEmpty(user.Lastname) {
		return User{}, ErrRequiredLastname
	}
	if IsEmpty(user.Email) {
		return User{}, ErrRequiredEmail
	}
	if !IsEmail(user.Email) {
		return User{}, ErrInvalidEmail
	}
	if IsEmpty(user.Password) {
		return User{}, ErrRequiredPassword
	}

	return user, nil
}
