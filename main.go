package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf(fmt.Sprintf("Server running on port:: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
