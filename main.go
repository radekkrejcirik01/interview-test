package main

import (
	"log"

	"interview_test/pkg/rest"
)

func main() {
	if err := rest.Create().Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
