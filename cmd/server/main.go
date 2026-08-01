package main

import (
	"log"

	"github.com/ArjunDev17/course-content-service/internal/app"
)

func main() {

	application, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}