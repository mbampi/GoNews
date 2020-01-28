package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"gonews/src/models"

	"github.com/gorilla/mux"
)

func (s *Server) getAllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: return all posts")

	rows, err := s.DB.Query("SELECT * FROM post")
	posts := []models.Post{}

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		post := models.Post{}
		rows.Scan(&post.ID, &post.Title, &post.Content)
		posts = append(posts, post)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func (s *Server) getPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: return single post")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	row := s.DB.QueryRow("SELECT * FROM post WHERE id=?", id)
	post := models.Post{}
	row.Scan(&post.ID, &post.Title, &post.Content)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func (s *Server) createPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: create new post")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var post models.Post
	json.Unmarshal(reqBody, &post)

	stmt, err := s.DB.Prepare("INSERT post SET id=?, title=?, content=?")
	_, err = stmt.Exec(post.ID, post.Title, post.Content)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("OK")
}

func (s *Server) updatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: update post")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	reqBody, _ := ioutil.ReadAll(r.Body)
	var post models.Post
	json.Unmarshal(reqBody, &post)
	fmt.Println(post)
	post.ID = id
	fmt.Println(post)

	stmt, err := s.DB.Prepare("UPDATE post SET id=?, title=?, content=? WHERE id=?")
	_, err = stmt.Exec(post.ID, post.Title, post.Content, id)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) deletePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: delete post")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	stmt, err := s.DB.Prepare("DELETE FROM post WHERE id=?")
	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("OK")
}

func (s *Server) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")

	fmt.Fprintf(w, "-- Welcome to my first Go API --")
	fmt.Fprintf(w, "\n Matheus D Bampi ")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("OK")
}
