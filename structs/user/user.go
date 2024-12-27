package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	// User // u can give it a name like user User
	User User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "John",
			lastName:  "Doe",
			birthdate: "01/01/2000",
			createdAt: time.Now(),
		},
	}
}

func (u User) OutputUserDetails() {
	fmt.Println("Struct method details")
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *User) ShowUserDetails() {
	fmt.Println("Struct method details:Pointer")
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// constructor function
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name is required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
