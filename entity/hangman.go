package entity

import (
	"fmt"
	"strings"
)

type Hangman struct {
	ID        int    `json:"id"`
	Word      string `json:"word"`
	Length    int    `json:"letters"`
	Remaining int    `json:"letters"`
	Guesses   int    `json:"guesses"`
	Status    string `json:"status"`
}

func (h Hangman) Guess(letter string) (Hangman, []int) {
	wordSlice := strings.Split(h.Word, "")

	var index []int

	i := 0
	found := 0
	for _, value := range wordSlice {
		if letter == value {
			index = append(index, i+1)
			found++
		}
		i++
	}

	h.Remaining -= found

	contain := strings.Contains(h.Word, letter)

	if !contain {
		h.Guesses -= 1
	}

	if !contain && h.Guesses == 0 {
		h.Status = "lost"
	}

	return h, index
}

var games []Hangman

func Store(hangman Hangman) {
	games = []Hangman{hangman}
}

func Find(id int) (Hangman, error) {
	for _, hangman := range games {
		if hangman.ID == id {
			return hangman, nil
		}
	}

	return Hangman{}, fmt.Errorf("game id: %v not found", id)
}
