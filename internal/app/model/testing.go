package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "test@gmail.com",
		Password: "password",
	}
}
