package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")

	sendEmail()

	fmt.Println("end")
}

func sendEmail() {
	time.Sleep(time.Second * 5)
}
