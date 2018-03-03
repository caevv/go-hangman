package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strconv"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hangman", CreateGame).Methods("POST")
	router.HandleFunc("/hangman", GetGames).Methods("GET")
	router.HandleFunc("/hangman/{id}", GetGame).Methods("GET")
	router.HandleFunc("/hangman/{id}/guess", Guess).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func Guess(response http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))

	body, _ := ioutil.ReadAll(r.Body)
	var bodyArray map[string]interface{}
	json.Unmarshal(body, &bodyArray)

	hangman, _ := Find(id)

	hangman, index := hangman.Guess(bodyArray["letter"].(string))

	respondWithJSON(response, 200, map[string]interface{}{
		"id": hangman.ID, "guesses": hangman.Guesses, "length": hangman.Length, "index": index, "status": hangman.Status,
	})
}

func GetGame(response http.ResponseWriter, request *http.Request) {
	//TODO
}
func GetGames(response http.ResponseWriter, request *http.Request) {
	//TODO
}

func CreateGame(w http.ResponseWriter, request *http.Request) {
	hangman := Hangman{
		ID:        1,
		Word:      "cryptocurrency",
		Length:    14,
		Guesses:   5,
		Remaining: 14,
		Status:    "ongoing",
	}

	Store(hangman)

	respondWithJSON(w, http.StatusCreated, map[string]int{"id": hangman.ID, "guesses": hangman.Guesses, "length": hangman.Length})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
