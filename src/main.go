package main

import (
	"log"
	"os"

	"gonews/src/controllers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	var s *controllers.Server
	s = new(controllers.Server)
	s.Initialize(dbUser, dbPassword, dbName)
	s.Run()
}
