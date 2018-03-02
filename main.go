package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hangman", CreateGame).Methods("POST")
	router.HandleFunc("/hangman", GetGames).Methods("GET")
	router.HandleFunc("/hangman/{id}", GetGame).Methods("GET")
	router.HandleFunc("/hangman/{id}/guess", Guess).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
func Guess(writer http.ResponseWriter, request *http.Request) {
	//TODO
}
func GetGame(writer http.ResponseWriter, request *http.Request) {
	//TODO
}
func GetGames(writer http.ResponseWriter, request *http.Request) {
	//TODO
}
func CreateGame(writer http.ResponseWriter, request *http.Request) {
	//TODO
}
