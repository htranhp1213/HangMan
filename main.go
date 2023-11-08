package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// short cut for print next line
var pl = fmt.Println

// list of words to guess
var wordList = []string{"programming", "hangman", "golang", "example", "developer"}

func main() {
	rand.Seed(time.Now().UnixNano())
	wordToGuess := getRandomWord()
	guesses := make([]string, 0)
	maxAttempts := 6 // maximum tries given by system
	attempts := 0    // number of tries taken by guess

	pl("You have %d attempts to guess the word.\n", maxAttempts)

	for {
		displayWord(wordToGuess, guesses)
		pl("Enter a letter:")
		guess := getUserInput()
		guesses = append(guesses, guess)

		/* if player input the letter which is not included in the given word
		increment their tries
		decrease the maximum tries
		*/
		if !strings.Contains(wordToGuess, guess) { // Contains is a built in func
			attempts++
			pl("Oops!!! Hang in there. Only /%d attempts left\n", maxAttempts-attempts)
		}

		if isWon(wordToGuess, guesses) {
			displayWord(wordToGuess, guesses)
			pl("Congratulations! You've won!")
			break
		}

		if maxAttempts == 0 {
			pl("No more attempts ^w^ Hang the man!!!")
		}
	}
}

func displayWord(wordToGuess string, guesses []string) {
	display := " "
	for _, letter := range wordToGuess {
		if contains(guesses, string(letter)) {
			display += string(letter)
		} else {
			display += "_"
		}
	}
	pl(display)
}

func isWon(wordToGuess string, guesses []string) bool {
	for _, letter := range wordToGuess {
		if !contains(guesses, string(letter)) {
			return false
		}
	}
	return true
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)   // reader works as scanner in Java
	guess, err := reader.ReadString('\n') // guess is where the user input stored, err is to protected guess from error
	if err != nil {                       // if err is not null
		log.Fatal(err)
	}
	return strings.TrimSpace(strings.ToLower(guess)) // if there is no error, return the guess with null error

}

func getRandomWord() string {
	return wordList[rand.Intn(len(wordList))]
}

// check if the guess is matched with the word to guess
func contains(wordToGuess []string, item string) bool {
	for _, letter := range wordToGuess {
		if letter == item {
			return true
		}
	}
	return false

}
