package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/mbampi/gonews/src/models"
)

// Post model
type Post = models.Post

// Author (fk of Post)
type Author = models.Author

var posts []Post

func getPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: return all posts")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
	return
}

func getPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: return single post")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, post := range posts {
		if post.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(post)
			return
		}
	}
}

func createPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: create new post")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var post Post
	json.Unmarshal(reqBody, &post)

	posts = append(posts, post)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
	return
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: create new post")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, item := range posts {
		if item.ID == id {
			posts = append(posts[:index], posts[index+1:]...)

			reqBody, _ := ioutil.ReadAll(r.Body)
			var post Post
			json.Unmarshal(reqBody, &post)
			posts = append(posts, post)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(post)
			return
		}
	}
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: delete post")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, post := range posts {
		if post.ID == id {
			posts = append(posts[:index], posts[index+1:]...)
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "-- Welcome to my first Go API --")
	fmt.Fprintf(w, "\n Matheus D Bampi ")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/post/{id}", getPost).Methods("GET")
	router.HandleFunc("/post", createPost).Methods("POST")
	router.HandleFunc("/post/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/post/{id}", deletePost).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	posts = []Post{
		Post{ID: 1, Title: "My first post", Content: "this is my first post", Author: &Author{CPF: 123, Name: "Matheus O Grande"}},
		Post{ID: 2, Title: "My second post", Content: "this is my second post", Author: &Author{CPF: 456, Name: "Bob Jorge"}},
	}

	handleRequests()
}
