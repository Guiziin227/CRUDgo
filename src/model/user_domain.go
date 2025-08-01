package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetUsername() string
	GetAge() int8

	SetID(string)

	GetJSONValue() (string, error)

	EncryptPassword()
}

// NewUserDomain creates a new instance of userDomain with the provided parameters.
func NewUserDomain(email, password, username string, age int8) *userDomain {
	return &userDomain{
		Email:    email,
		Password: password,
		Username: username,
		Age:      age,
	}
}

func (ud *userDomain) SetID(id string) {
	ud.ID = id
}

// userDomain represents the user domain model with methods to manipulate user data.
type userDomain struct {
	ID       string
	Email    string
	Password string
	Username string
	Age      int8
}

func (ud *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		println(err.Error())
		return "", err
	}
	return string(b), nil
}

// GetEmail returns the email of the user.
func (ud *userDomain) GetEmail() string {
	return ud.Email
}

// GetPassword returns the password of the user.
func (ud *userDomain) GetPassword() string {
	return ud.Password
}

// GetUsername returns the username of the user.
func (ud *userDomain) GetUsername() string {
	return ud.Username
}

// GetAge returns the age of the user.
func (ud *userDomain) GetAge() int8 {
	return ud.Age
}

// EncryptPassword hashes the user's password using MD5.
func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
