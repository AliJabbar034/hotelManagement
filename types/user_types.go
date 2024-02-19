package types

import (
	"errors"
	"fmt"
	"strings"
)

type User struct {
	ID        string `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string `json:"first_name,omitempty" bson="first_name,omitempty"`
	LastName  string `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Email     string `json:"email,omitempty" bson:"email,omitempty"`
	Password  string `json:"password" bson:"password,omitempty"`
	Role      string `json:"role,omitempty" bson:"role,omitempty"`
}

func NewUser(user User) *User {
	return &User{

		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		Role:      "user",
	}
}

func ValidateInput(user User) error {
	fmt.Println(user)
	if user.FirstName == "" || user.LastName == "" || len(user.Password) < 6 || !strings.Contains("@", user.Email) {
		fmt.Println(strings.Contains("@", user.Email))
		return errors.New("please fill in all fields")
	}

	return nil
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ValidateLoginData(loginData LoginData) error {
	// res := strings.Contains("@", loginData.Email)

	if len(loginData.Password) < 6 {

		return errors.New("Invalid login data")

	}
	return nil
}
