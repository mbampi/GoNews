package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/mbampi/gonews/src/controllers"
	"github.com/mbampi/gonews/src/models"
)

// Post model
type Post = models.Post

// Author (fk of Post)
type Author = models.Author

var posts []Post

func main() {
	// posts = []Post{
	// 	Post{ID: 1, Title: "My first post", Content: "this is my first post", Author: &Author{CPF: 123, Name: "Matheus O Grande"}},
	// 	Post{ID: 2, Title: "My second post", Content: "this is my second post", Author: &Author{CPF: 456, Name: "Bob Jorge"}},
	// }

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
