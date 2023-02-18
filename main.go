package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gonews/controllers"
	"gonews/models"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func initDatabase() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	var err error
	controllers.DB, err = gorm.Open("mysql", dbUser+":"+dbPassword+"@/"+dbName+"?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer controllers.DB.Close()

	controllers.DB.AutoMigrate(&models.Post{})
	controllers.DB.AutoMigrate(&models.Author{})
}

func setupRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controllers.HomePage)
	router.HandleFunc("/posts", controllers.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", controllers.GetPost).Methods("GET")
	router.HandleFunc("/posts", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", controllers.DeletePost).Methods("DELETE")

	router.HandleFunc("/authors", controllers.GetAllAuthors).Methods("GET")
	router.HandleFunc("/authors/{id}", controllers.GetAuthor).Methods("GET")
	router.HandleFunc("/authors", controllers.CreateAuthor).Methods("POST")
	router.HandleFunc("/authors/{id}", controllers.UpdateAuthor).Methods("PUT")
	router.HandleFunc("/authors/{id}", controllers.DeleteAuthor).Methods("DELETE")

	return router
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	initDatabase()

	router := setupRoutes()

	fmt.Println("Running")
	log.Fatal(http.ListenAndServe(":8081", router))
}
