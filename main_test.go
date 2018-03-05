package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/godog"
	"net/http"
	"net/http/httptest"

	"go-hangman-api/entity"
)

type apiFeature struct {
	resp *httptest.ResponseRecorder
}

func (a *apiFeature) resetResponse(interface{}) {
	a.resp = httptest.NewRecorder()
}

func (a *apiFeature) iStartANewGame() (err error) {
	req, err := http.NewRequest("POST", "/hangman", nil)

	if err != nil {
		return err
	}

	CreateGame(a.resp, req)

	return
}

func (a *apiFeature) thereShouldBeAGameWithWordWithRemainingGuesses(word string, guesses int) error {
	type hangmanResponse struct {
		ID int `json:"id"`
	}

	var r hangmanResponse
	json.Unmarshal(a.resp.Body.Bytes(), &r)

	hanghman, err := entity.Find(r.ID)

	if err != nil {
		return err
	}

	if word != hanghman.Word {
		return fmt.Errorf("expected word to be: %s, but actual is: %s", word, hanghman.Word)
	}

	if guesses != hanghman.Guesses {
		return fmt.Errorf("expected guesses to be: %v, but actual is: %v", word, hanghman.Word)
	}

	return nil
}

func (a *apiFeature) iShouldBeToldThatTheWordHaveLettersAndRemainingGuessesWitAnID(length int, guesses int) error {
	if a.resp.Code != 201 {
		return fmt.Errorf("expected response code to be: %d, but actual is: %#v", 201, a.resp.Code)
	}

	type hangmanResponse struct {
		ID      *int `json:"id"`
		Guesses int  `json:"guesses"`
		Length  int  `json:"length"`
	}

	var hangman hangmanResponse

	json.Unmarshal(a.resp.Body.Bytes(), &hangman)

	if hangman.ID == nil {
		return fmt.Errorf("expected id to be an integer but actual is: %v", hangman.ID)
	}

	if guesses != hangman.Guesses {
		return fmt.Errorf("expected guesses to be: %v, but actual is: %v", guesses, hangman.Guesses)
	}

	if length != hangman.Length {
		return fmt.Errorf("expected guesses to be: %v, but actual is: %v", length, hangman.Length)
	}

	return nil
}

func (a *apiFeature) thereIsAGameStartedWithWordWithRemainingGuessesWithLetters(word string, guesses int, length int) error {
	hangmam := entity.Hangman{
		ID:      1,
		Word:    word,
		Length:  length,
		Guesses: guesses,
	}

	entity.Store(hangmam)

	return nil
}

func (a *apiFeature) iGuessAGuessTheLetter(letter string) error {
	payload := []byte(fmt.Sprintf(`{"letter":"%s"}`, letter))

	req, err := http.NewRequest("POST", "/hangman/1/", bytes.NewBuffer(payload))

	if err != nil {
		return err
	}

	Guess(a.resp, req)

	return nil
}

func (a *apiFeature) iShouldBeToldTheLetterIsWrong() error {
	if a.resp.Code != 200 {
		return fmt.Errorf("expected response code to be: %d, but actual is: %#v", 200, a.resp.Code)
	}

	type hangmanResponse struct {
		Index []int `json:"index"`
	}

	var hangman hangmanResponse

	json.Unmarshal(a.resp.Body.Bytes(), &hangman)

	if len(hangman.Index) > 0 {
		return fmt.Errorf("expected not words to be found but found: %v", len(hangman.Index))
	}

	return nil
}

func (a *apiFeature) thatIHaveOtherAttempts(guesses int) error {
	type hangmanResponse struct {
		Guesses int `json:"guesses"`
	}

	var hangman hangmanResponse

	json.Unmarshal(a.resp.Body.Bytes(), &hangman)

	if hangman.Guesses != guesses {
		return fmt.Errorf("expected guesses to be: %d, but actual is: %#v", guesses, hangman.Guesses)
	}

	return nil
}

func (a *apiFeature) theIHaveTriedLetter(letter string) error {
	type hangmanResponse struct {
		Letters []string `json:"letters"`
	}

	var hangman hangmanResponse

	json.Unmarshal(a.resp.Body.Bytes(), &hangman)

	found := false
	for _, value := range hangman.Letters {
		if letter == value {
			found = true
		}
	}


	if !found {
		return fmt.Errorf("expected letter %s to be found, but it was not", letter)
	}

	return nil
}

func (a *apiFeature) iShouldBeToldTheLetterIsAvailableOnPositionsAnd(index1, index2, index3 int) error {
	if a.resp.Code != 200 {
		return fmt.Errorf("expected response code to be: %d, but actual is: %#v", 200, a.resp.Code)
	}

	type hangmanResponse struct {
		Index []int `json:"index"`
	}

	var hangman hangmanResponse

	json.Unmarshal(a.resp.Body.Bytes(), &hangman)

	if len(hangman.Index) != 3 {
		return fmt.Errorf("expected 3 indeces to be found, but found: %v", len(hangman.Index))
	}

	var found1 bool
	for _, i := range hangman.Index {
		if i == index1 {
			found1 = true
		}
	}
	var found2 bool
	for _, i := range hangman.Index {
		if i == index2 {
			found2 = true
		}
	}
	var found3 bool
	for _, i := range hangman.Index {
		if i == index3 {
			found3 = true
		}
	}

	if !found1 {
		return fmt.Errorf("expected index %d to be found but was not", index1)
	}

	if !found2 {
		return fmt.Errorf("expected index %d to be found but was not", index2)
	}

	if !found3 {
		return fmt.Errorf("expected index %d to be found but was not", index3)
	}

	return nil
}

func (a *apiFeature) thatILostTheGame() error {
	type hangmanResponse struct {
		Status string `json:"status"`
	}

	var hangman hangmanResponse

	json.Unmarshal(a.resp.Body.Bytes(), &hangman)

	if hangman.Status != "lost" {
		return fmt.Errorf("expected to lost the game, but found: %s", hangman.Status)
	}

	return nil
}

func (a *apiFeature) thereIsATerminatedGame() error {
	return godog.ErrPending
}

func (a *apiFeature) iGuessSomeLetter() error {
	return godog.ErrPending
}

func (a *apiFeature) iShouldBeToldTheGameHasEnded() error {
	return godog.ErrPending
}

func (a *apiFeature) thereIsAWonGameWithWord(arg1 string) error {
	return godog.ErrPending
}

func (a *apiFeature) iListTheGames() error {
	return godog.ErrPending
}

func (a *apiFeature) iShouldBeThatThereIsAGameWithWordWithRemainingGuesses(arg1, arg2 string) error {
	return godog.ErrPending
}

func (a *apiFeature) iShouldBeToldThatThereIsAWonGameWithWord(arg1 string) error {
	return godog.ErrPending
}

func (a *apiFeature) serverIsNotAvailable() error {
	return godog.ErrPending
}

func (a *apiFeature) iShouldBeToldThatThereIsAServerError() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	api := &apiFeature{}

	s.BeforeScenario(api.resetResponse)

	s.Step(`^I start a new game$`, api.iStartANewGame)
	s.Step(`^there should be a game with word "([^"]*)" with "([^"]*)" remaining guesses$`, api.thereShouldBeAGameWithWordWithRemainingGuesses)
	s.Step(`^I should be told that the word have "([^"]*)" letters and "([^"]*)" remaining guesses wit an ID$`, api.iShouldBeToldThatTheWordHaveLettersAndRemainingGuessesWitAnID)
	s.Step(`^there is a game started with word "([^"]*)" with "([^"]*)" remaining guesses with "([^"]*)" letters$`, api.thereIsAGameStartedWithWordWithRemainingGuessesWithLetters)
	s.Step(`^the I have tried letter "([^"]*)"$`, api.theIHaveTriedLetter)
	s.Step(`^I guess a guess the letter "([^"]*)"$`, api.iGuessAGuessTheLetter)
	s.Step(`^I should be told the letter is wrong$`, api.iShouldBeToldTheLetterIsWrong)
	s.Step(`^that I have other "([^"]*)" attempts$`, api.thatIHaveOtherAttempts)
	s.Step(`^I should be told the letter is available on positions "([^"]*)", "([^"]*)" and "([^"]*)"$`, api.iShouldBeToldTheLetterIsAvailableOnPositionsAnd)
	s.Step(`^that I lost the game$`, api.thatILostTheGame)
	s.Step(`^there is a terminated game$`, api.thereIsATerminatedGame)
	s.Step(`^I guess some letter$`, api.iGuessSomeLetter)
	s.Step(`^I should be told the game has ended$`, api.iShouldBeToldTheGameHasEnded)
	s.Step(`^there is a won game with word "([^"]*)"$`, api.thereIsAWonGameWithWord)
	s.Step(`^I list the games$`, api.iListTheGames)
	s.Step(`^I should be told that there is a game with word "([^"]*)" with "([^"]*)" remaining guesses$`, api.iShouldBeThatThereIsAGameWithWordWithRemainingGuesses)
	s.Step(`^I should be told that there is a won game with word "([^"]*)"$`, api.iShouldBeToldThatThereIsAWonGameWithWord)
	s.Step(`^server is not available$`, api.serverIsNotAvailable)
	s.Step(`^I should be told that there is a server error$`, api.iShouldBeToldThatThereIsAServerError)
}
