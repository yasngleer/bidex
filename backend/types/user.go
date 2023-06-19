package types

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int    `bson:"_id,omitempty" json:"id,omitempty"`
	Email             string `bson:"email" json:"email,omitempty"`
	EncryptedPassword string `bson:"encryptedPassword" json:"_,omitempty"`
}

type UserResponse struct {
	ID    int    `bson:"_id,omitempty" json:"id,omitempty"`
	Email string `bson:"email" json:"email,omitempty"`
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

func (u User) ToResponse() UserResponse {
	return UserResponse{
		Email: u.Email,
		ID:    u.ID,
	}
}
