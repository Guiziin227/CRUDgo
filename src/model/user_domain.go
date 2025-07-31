package model

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/guiziin227/CRUDgo/src/configuration/c_err"
)

func NewUserDomain(email, password, username string, age int8) *userDomain {
	return &userDomain{
		email:    email,
		password: password,
		username: username,
		age:      age,
	}
}

type userDomain struct {
	email    string
	password string
	username string
	age      int8
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *c_err.CErr
	UpdateUser(string) *c_err.CErr
	FindUser(string) (*userDomain, *c_err.CErr)
	DeleteUser(string) *c_err.CErr
}
