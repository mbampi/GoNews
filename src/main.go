package main

import (
	"log"
	"os"

	"gonews/src/controllers"
	"gonews/src/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Post model
type Post = models.Post

// Author (fk of Post)
type Author = models.Author

var posts []Post

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
