package main

import "fmt"

func main() {

	// Familiar "for"
	for i := 0; i < 5; i++ {
		fmt.Println("Value of i is now:", i)
	}

	// Same as a "while" in php, ie. loop forever
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Println("Value of i is:", i)
		i++
	}

	// Same as a "foreach" in php
	a := [...]string{"a", "b", "c", "d"}
	for i, v := range a {
		fmt.Println("Item", i, "is", v)
	}

	// Switch - pretty familiar to php
	s := "bob"
	switch s {
	case "fred", "frank", "mary":
		fmt.Println("Not bob")
	case "bob":
		fmt.Println("Is bob")
	}
}
