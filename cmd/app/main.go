package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	env := os.Getenv("env")
	if err != nil && env == "" {
		log.Fatalf("Failed to load env, err : %v", err)
	}

	repository := repository.NewRepository(db)

	service := service.NewService(repository)

	middleware := middleware.Init(service)

	rest := rest.NewRest(service, middleware)

}
