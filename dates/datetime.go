package main

import (
	"fmt"
	"time"
)

func main() {

	format := "Mon Jan 2 15:04:05 MST 2006  (MST is GMT-0700)"

	t := time.Now()
	fmt.Println(t.Format(format))
}
