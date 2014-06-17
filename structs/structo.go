package main

import "fmt"

/**
 * NOTE: if you haven't noticed yet, go is not an object oriented language
 */
type User struct {
	FirstName string
	SurName   string
	Age       int
	Email     string
}

func (u *User) FullName() string {
	return u.FirstName + " " + u.SurName
}

func main() {
	u := User{"Bob", "Smith", 33, "bob@bob.com"}

	usersFirstName := u.FirstName

	usersFullName := u.FullName()

	fmt.Println("First name: ", usersFirstName)
	fmt.Println("FULL name: ", usersFullName)

	u.Save()

}

func (u *User) Save() error {
	// Save user to the database
	return nil
}
