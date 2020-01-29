package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Server contains Database and Router
type Server struct {
	DB     *sql.DB
	Router *mux.Router
}

// Initialize the server
func (s *Server) Initialize(dbUser, dbPassword, dbName string) {
	var err error
	s.DB, err = sql.Open("mysql", dbUser+":"+dbPassword+"@/"+dbName)
	if err != nil {
		log.Fatal(err)
	}

	s.Router = mux.NewRouter().StrictSlash(true)
	s.initializeRouter()
}

// Run listening to port defined
func (s *Server) Run() {
	log.Fatal(http.ListenAndServe(":8081", s.Router))
	fmt.Println("Running")
}

func (s *Server) initializeRouter() {
	s.Router.HandleFunc("/", s.homePage)
	s.Router.HandleFunc("/posts", s.getAllPosts).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", s.getPost).Methods("GET")
	s.Router.HandleFunc("/posts", s.createPost).Methods("POST")
	s.Router.HandleFunc("/posts/{id}", s.updatePost).Methods("PUT")
	s.Router.HandleFunc("/posts/{id}", s.deletePost).Methods("DELETE")
}
