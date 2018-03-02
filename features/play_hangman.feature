Feature: Play Hangman
    As a user I want to be able to start a new game;
    As a user I want to be able to guess a character for an ongoing game;
    As a user I want to be notified when the game I am playing ends (win/game over);
    As a user I want to resume an incomplete game;
    As a user I want to be able to list all the games that have been played so far and the ones that are currently ongoing;
    As a user I want to be notified when I trigger an action which results in a server failure.

  Scenario: Start a new game
    When I start a new game
    Then I should be told that the word have "14" letters and "5" remaining guesses

  Scenario: Wrong guess
    Given there is a game started with word "cryptocurrency" with "5" remaining guesses
    When I guess a guess the letter "a"
    Then I should be told the letter is wrong
    And that I have other "4" attempts

  Scenario: Correct guess
    Given there is a game started with word "cryptocurrency" with "4" remaining guesses
    When I guess a guess the letter "c"
    Then I should be told the letter is correcy as "c_____c_____c_"
    And that I have other "4" attempts

  Scenario: Lost Game
    Given there is a game started with word "cryptocurrency" with "1" remaining guesses
    When I guess a guess the letter "z"
    Then I should be told the letter is wrong
    And that I lost the game

  Scenario: Ended game non available
    Given there is a terminated game
    When I guess some letter
    Then I should be told the game has ended

  Scenario: List games
    Given there is a game started with word "cryptocurrency" with "5" remaining guesses
    And there is a won game with word "cryptocurrency"
    When I list the games
    Then I should be that there is a game with word "cryptocurrency" with "5" remaining guesses
    And I should be told that there is a won game with word "cryptocurrency"

  Scenario: Server error
    Given server is not available
    When I start a new game
    Then I should be told that there is a server error
