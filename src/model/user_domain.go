package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetUsername() string
	GetAge() int8

	EncryptPassword()
}

// NewUserDomain creates a new instance of userDomain with the provided parameters.
func NewUserDomain(email, password, username string, age int8) *userDomain {
	return &userDomain{
		email:    email,
		password: password,
		username: username,
		age:      age,
	}
}

// userDomain represents the user domain model with methods to manipulate user data.
type userDomain struct {
	email    string
	password string
	username string
	age      int8
}

// GetEmail returns the email of the user.
func (ud *userDomain) GetEmail() string {
	return ud.email
}

// GetPassword returns the password of the user.
func (ud *userDomain) GetPassword() string {
	return ud.password
}

// GetUsername returns the username of the user.
func (ud *userDomain) GetUsername() string {
	return ud.username
}

// GetAge returns the age of the user.
func (ud *userDomain) GetAge() int8 {
	return ud.age
}

// EncryptPassword hashes the user's password using MD5.
func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}
