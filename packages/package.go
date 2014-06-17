package model

import "fmt"

type User struct {
	Id    int
	Name  string
	Email string
}

var NumberOfUsers int = 2

func GetUsers() []User {

	users := []User{User{1, "Andrew", "ae@ae.ae"}, User{2, "bob", "bob@bob.bob"}}

	return users
}

func onlyAvailableInPackage() {
	fmt.Println("Can't access this outside of the 'model' package")
}
