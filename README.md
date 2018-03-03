# Go Hangman API

## Testing

- run `make test`.
- BDD tests were created with Godog. 


```gherkin
Feature: Play Hangman
  Scenario: Start a new game                                                                 
    When I start a new game                                                                  
    Then there should be a game with word "cryptocurrency" with "5" remaining guesses        
    And I should be told that the word have "14" letters and "5" remaining guesses wit an ID 

  Scenario: Wrong guess                                                                                   
    Given there is a game started with word "cryptocurrency" with "5" remaining guesses with "14" letters 
    When I guess a guess the letter "z"                                                                   
    Then I should be told the letter is wrong                                                             
    And that I have other "4" attempts                                                                    

  Scenario: Correct guess                                                                                 
    Given there is a game started with word "cryptocurrency" with "5" remaining guesses with "14" letters 
    When I guess a guess the letter "c"                                                                   
    Then I should be told the letter is available on positions "1", "7" and "13"                          
    And that I have other "5" attempts                                                                    

  Scenario: Lost Game                                                                                     
    Given there is a game started with word "cryptocurrency" with "1" remaining guesses with "14" letters 
    When I guess a guess the letter "z"                                                                   
    Then I should be told the letter is wrong                                                             
    And that I lost the game                                                                              

4 scenarios (4 passed)
15 steps (15 passed)

```


### TODO:
- Move some logics from main to handler, so they would have just struct as parameter.
- Store letters already attempted.
- Check if won.
- List games.
- Get a game status.
- Fail gracefully.
