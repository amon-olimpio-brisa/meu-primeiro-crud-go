package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8

	EncryptPassword()
}

func NewUserDomain(
	email, Password, name string,
	age int8,
) UserDomainInterface {
	return &UserDomain{email, Password, name, age}
}

type UserDomain struct {
	email    string
	password string
	name     string
	age      int8
}

func (ud *UserDomain) GetEmail() string {
	return ud.email
}

func (ud *UserDomain) GetPassword() string {
	return ud.password
}

func (ud *UserDomain) GetName() string {
	return ud.name
}

func (ud *UserDomain) GetAge() int8 {
	return ud.age
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}
