package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gonews/api/controllers"
	"gonews/api/models"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
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

	controllers.DB, err = gorm.Open("mysql", dbUser+":"+dbPassword+"@/"+dbName+"?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer controllers.DB.Close()

	controllers.DB.AutoMigrate(&models.Post{})
	controllers.DB.AutoMigrate(&models.Author{})

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controllers.HomePage)
	router.HandleFunc("/posts", controllers.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", controllers.GetPost).Methods("GET")
	router.HandleFunc("/posts", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", controllers.DeletePost).Methods("DELETE")

	fmt.Println("Running")
	log.Fatal(http.ListenAndServe(":8081", router))
}