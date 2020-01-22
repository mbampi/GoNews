package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mbampi/gonews/src/models"
)

func (s *Server) getPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: return all posts")

	posts, err := getAllPosts(s.DB)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
	return
}

func getAllPosts(db *sql.DB) ([]models.Post, error) {
	rows, err := db.Query("SELECT * FROM post")
	posts := []models.Post{}

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		post := models.Post{}
		rows.Scan(&post.ID, &post.Title, &post.Content)
		posts = append(posts, post)
		fmt.Println("Post => ID:" + strconv.Itoa(post.ID) + " | Title:" + post.Title + " | Content:" + post.Content)
	}
	return posts, err
}

func (s *Server) getPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: return single post")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	print(id)
	// for _, post := range posts {
	// 	if post.ID == id {
	// 		w.Header().Set("Content-Type", "application/json")
	// 		json.NewEncoder(w).Encode(post)
	// 		return
	// 	}
	// }
}

func (s *Server) createPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: create new post")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var post models.Post
	json.Unmarshal(reqBody, &post)

	// posts = append(posts, post)

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(post)
	// return
}

func (s *Server) updatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: create new post")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	print(id)
	// for index, item := range posts {
	// 	if item.ID == id {
	// 		posts = append(posts[:index], posts[index+1:]...)

	// 		reqBody, _ := ioutil.ReadAll(r.Body)
	// 		var post Post
	// 		json.Unmarshal(reqBody, &post)
	// 		posts = append(posts, post)

	// 		w.Header().Set("Content-Type", "application/json")
	// 		json.NewEncoder(w).Encode(post)
	// 		return
	// 	}
	// }
}

func (s *Server) deletePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: delete post")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	print(id)
	// for index, post := range posts {
	// 	if post.ID == id {
	// 		posts = append(posts[:index], posts[index+1:]...)
	// 	}
	// }
}

func (s *Server) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "-- Welcome to my first Go API --")
	fmt.Fprintf(w, "\n Matheus D Bampi ")
	fmt.Println("Endpoint Hit: homePage")
}
