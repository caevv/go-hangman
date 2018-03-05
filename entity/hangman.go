package entity

import (
	"fmt"
	"strings"
)

type Hangman struct {
	ID        int      `json:"id"`
	Word      string   `json:"word"`
	Length    int      `json:"letters"`
	Remaining int      `json:"remaining"` // remaining letters
	Guesses   int      `json:"guesses"`	  // remaining guesses
	Status    string   `json:"status"`
	Letters   []string `json:"letters"`
}

func (h Hangman) Guess(letter string) (Hangman, []int) {
	h.Letters = append(h.Letters, letter)

	// search letter on the word
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

	if found == 0 {
		h.Guesses -= 1
	}

	if found == 0 && h.Guesses == 0 {
		h.Status = "lost"
	}

	update(h)

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

func update(h Hangman) {
	for i, hangman := range games {
		if hangman.ID == h.ID {
			games[i] = h
		}
	}
}
