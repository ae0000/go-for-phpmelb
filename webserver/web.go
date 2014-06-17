package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Looking for %s", r.URL.Path[1:])
}

func main() {

	fmt.Println("listening: http://localhost:8080")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
