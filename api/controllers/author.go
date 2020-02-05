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

// GetAllAuthors from given database
func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: return all authors")

	var authors []models.Author
	result := DB.Find(&authors)
	if result.Error != nil {
		responses.RespondWithError(w, http.StatusBadRequest, result.Error)
		return
	}

	responses.RespondWithJSON(w, http.StatusOK, authors)
}

// GetAuthor handles to return the author referent to the id
func GetAuthor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: return single author")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var author models.Author
	result := DB.First(&author, id)
	if result.Error != nil {
		responses.RespondWithError(w, http.StatusBadRequest, result.Error)
		return
	}

	responses.RespondWithJSON(w, http.StatusOK, author)
}

// CreateAuthor adds to database a new author
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: create new author")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var author models.Author
	json.Unmarshal(reqBody, &author)

	err := author.Validate()
	if err != nil {
		responses.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}

	result := DB.Create(&author)
	if result.Error != nil {
		log.Println(result.Error)
		responses.RespondWithError(w, http.StatusBadRequest, result.Error)
		return
	}

	responses.RespondWithJSON(w, http.StatusOK, author)
}

// UpdateAuthor updates author
func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: update author")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	reqBody, _ := ioutil.ReadAll(r.Body)
	var author models.Author
	json.Unmarshal(reqBody, &author)

	err := author.Validate()
	if (err != nil){
		responses.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	author.ID = uint64(id)
	result := DB.Save(&author)
	if result.Error != nil {
		responses.RespondWithError(w, http.StatusBadRequest, result.Error)
		return
	}
}

// DeleteAuthor is
func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: delete author")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var author models.Author
	if id >= 0 {
		author.ID = uint64(id)
		result := DB.Delete(&author)
		if result.Error != nil {
			responses.RespondWithError(w, http.StatusBadRequest, result.Error)
			return
		}
	}

	responses.RespondWithJSON(w, http.StatusOK, "OK")
}