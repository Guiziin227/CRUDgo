package model

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/guiziin227/CRUDgo/src/configuration/c_err"
)

func NewUserDomain(email, password, username string, age int8) *UserDomain {
	return &UserDomain{
		Email:    email,
		Password: password,
		Username: username,
		Age:      age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Username string
	Age      int8
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *c_err.CErr
	UpdateUser(string) *c_err.CErr
	FindUser(string) (*UserDomain, *c_err.CErr)
	DeleteUser(string) *c_err.CErr
}
