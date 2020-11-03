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
	ErrMaxLimit          = errors.New("Greater than maximum number of characters")
	ErrDuplicateKeyEmail = errors.New("User with such email already exists")
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

func Max(attr string, lim int) bool {
	if len(attr) <= lim {
		return true
	}

	return false
}

func ValidateLimitFields(user User) (User, error) {
	if !Max(user.Firstname, 15) || !Max(user.Lastname, 20) || !Max(user.Email, 40) || !Max(user.Password, 10) {
		return user, ErrMaxLimit
	}
	return user, nil
}

func UniqueEmail(email string) (bool, error) {
	con := Connect()
	defer con.Close()

	sql := "select count(email) from users where email = $1"

	rs, err := con.Query(sql, email)
	if err != nil {
		return false, err
	}
	defer rs.Close()

	var count int64

	if rs.Next() {
		err := rs.Scan(&count)
		if err != nil {
			return false, err
		}
	}

	if count > 0 {
		return false, ErrDuplicateKeyEmail
	}

	return true, nil
}

func ValidateNewUser(user User) (User, error) {
	_, err := UniqueEmail(user.Email)
	if err != nil {
		return User{}, err
	}

	user, err = ValidateLimitFields(user)
	if err != nil {
		return user, err
	}

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
