package types

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                string `bson:"_id,omitempty" json:"id,omitempty"`
	Email             string `bson:"email" json:"email,omitempty"`
	EncryptedPassword string `bson:"encryptedPassword" json:"_,omitempty"`
}

type UserResponse struct {
	Email string `json:"email,omitempty"`
	Token string `json:"token,omitempty"`
}

func NewUser(email, password string) (*User, error) {
	epw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		Email:             email,
		EncryptedPassword: string(epw),
	}, nil
}

func (u *User) ValidatePassword(pw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(pw))
	return err == nil
}
