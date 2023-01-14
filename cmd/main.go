package main

import (
	"log"

	"github.com/bypepe77/secret-app-backend/cmd/bootstrap"
)

func main() {
	err := bootstrap.Run()
	if err != nil {
		log.Fatal(err)
	}
}
