package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"gonews/api/models"
	"gonews/api/responses"

	"github.com/gorilla/mux"
)

// GetAllPosts from given database
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: return all posts")

	var posts []models.Post
	result := DB.Preload("Author").Find(&posts)
	if result.Error != nil {
		responses.RespondWithError(w, http.StatusBadRequest, result.Error)
		return
	}

	responses.RespondWithJSON(w, http.StatusOK, posts)
}

// GetPost handles to return the post referent to the id
func GetPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: return single post")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var post models.Post
	result := DB.Preload("Author").First(&post, id)
	if result.Error != nil {
		responses.RespondWithError(w, http.StatusBadRequest, result.Error)
		return
	}

	responses.RespondWithJSON(w, http.StatusOK, post)
}

// CreatePost adds to database a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: create new post")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var post models.Post
	json.Unmarshal(reqBody, &post)

	err := post.Validate()
	if err != nil {
		responses.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}

	result := DB.Create(&post)
	if result.Error != nil {
		log.Println(result.Error)
		responses.RespondWithError(w, http.StatusBadRequest, result.Error)
		return
	}

	responses.RespondWithJSON(w, http.StatusOK, post)
}

// UpdatePost updates post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: update post")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	reqBody, _ := ioutil.ReadAll(r.Body)
	var post models.Post
	json.Unmarshal(reqBody, &post)

	err := post.Validate()
	if (err != nil){
		responses.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	post.ID = uint64(id)
	result := DB.Save(&post)
	if result.Error != nil {
		responses.RespondWithError(w, http.StatusBadRequest, result.Error)
		return
	}
}

// DeletePost is
func DeletePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: delete post")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var post models.Post
	if id >= 0 {
		post.ID = uint64(id)
		result := DB.Delete(&post)
		if result.Error != nil {
			responses.RespondWithError(w, http.StatusBadRequest, result.Error)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// HomePage prints an welcome message
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")

	fmt.Fprintf(w, "-- Welcome to my first Go API --")
	fmt.Fprintf(w, "\n Matheus D Bampi ")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("OK")
}
