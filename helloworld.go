package main

import (
	"errors"
	"fmt"
)

func main() {

	s := "Hello world"

	fmt.Println(s)

	s, i, err := returnMany()

	if err != nil {
		panic(err)
	}

	fmt.Println(s, i, err)
}

func returnMany() (string, int, error) {

	e := errors.New("Oh no")
	return "foo", 123, e
}
