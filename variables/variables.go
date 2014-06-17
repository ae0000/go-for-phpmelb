package main

import "fmt"

func main() {

	var a string = "bob"
	var b int = 123
	var c bool = true

	m := map[string]string{"Name": "bob", "Email": "bob@bob.bob"}

	fmt.Println(a, b, c)
}

// func returnMany() (string, int, bool) {
// 	return "foo", 123, true
// }

// func returnAnErrorMaybe() error {
// 	// Set the random seed and get a random number
// 	rand.Seed(time.Now().Unix())
// 	r := rand.Intn(3) + 1

// 	// 2 is an error
// 	if r == 2 {
// 		return errors.New("Oh no")
// 	} else {
// 		return nil
// 	}
// }
