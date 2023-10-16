package main

import (
	"log"

	"interview_test/pkg/rest"
)

func main() {
	// Run HTTP server on port 8080
	if err := rest.Create().Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
