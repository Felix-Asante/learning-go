package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var appUser User = User{
	// 	firstName,
	// 	lastName,
	// 	birthdate,
	// 	time.Now(),
	// }
	// var appUser user.User = User{
	// 	firstName: firstName,
	// 	lastName:  lastName,
	// 	birthdate: birthdate,
	// 	createdAt: time.Now(),
	// }

	// outputUserDetails(appUser)
	// outputUserDetails(&appUser)

	//METHODS
	// var user = User{
	// 	firstName: "John",
	// 	lastName:  "Doe",
	// 	birthdate: "01/01/2000",
	// 	createdAt: time.Now(),
	// }
	// user.outputUserDetails()
	// user.showUserDetails()
	// user.clearUserName()
	// user.showUserDetails()

	newUser, error := user.New(firstName, lastName, birthdate)
	if error != nil {
		fmt.Println(error)
		return
	}
	newUser.OutputUserDetails()
	newAdmin := user.NewAdmin("Rzv0Q@example.com", "password")
	newAdmin.User.OutputUserDetails()
}

// func outputUserDetails(user *User) {
// 	u := *user
// 	fmt.Println("User Details")
// 	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
// }

// func outputUserDetails(user *User) {
// 	fmt.Println(user.firstName, user.lastName, user.birthdate, user.createdAt)
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
