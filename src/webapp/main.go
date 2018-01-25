package main

import (
	"fmt"
	"net/http"
	"webapp/controller"
)

func main() {
	//todo populate templates here

	// Setup controller
	controller.Setup()

	fmt.Println("Server listening on 8080...")
	http.ListenAndServe(":8080", nil)
}
