package main

import (
	"fmt"
	"net/http"

	"gitlab.nordstrom.com/online-booking/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf(fmt.Sprintf("Server running on port:: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
