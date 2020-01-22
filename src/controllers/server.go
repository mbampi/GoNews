package controllers

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Server struct {
	DB     *sql.DB
	Router *mux.Router
}

func (s *Server) Initialize(dbUser, dbPassword, dbName string) {
	var err error
	s.DB, err = connectDB(dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatal(err)
	}

	s.Router = createRouter()
	s.initializeRouter()
}

func (s *Server) Run() {
	log.Fatal(http.ListenAndServe(":8081", s.Router))
}

func connectDB(dbUser, dbPassword, dbName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbUser+":"+dbPassword+"@/"+dbName)
	//defer db.Close()

	return db, err
}

func createRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	return router
}

func (s *Server) initializeRouter() {

	s.Router.HandleFunc("/", s.homePage)
	s.Router.HandleFunc("/posts", s.getPosts).Methods("GET")
	s.Router.HandleFunc("/post/{id}", s.getPost).Methods("GET")
	s.Router.HandleFunc("/post", s.createPost).Methods("POST")
	s.Router.HandleFunc("/post/{id}", s.updatePost).Methods("PUT")
	s.Router.HandleFunc("/post/{id}", s.deletePost).Methods("DELETE")
}
