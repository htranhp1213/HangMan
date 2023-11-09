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
		
		if attempts == maxAttempts {
			pl("No more attempts ^w^ Hang the man!!!")

			//Ask if the player wants to play again
			pl("Do you want to play again? (yes/no)")
			playAgain := getUserInput()
			if playAgain != "yes" {
				pl("Thanks for playing!")
				break
			} else {
				// Reset game variables for a new game
				wordToGuess = getRandomWord()
				guesses = make([]string, 0)
				attempts = 0
				pl("You have %d attempts to guess the word.\n", maxAttempts)
			}
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
// isLetter checks if the input is a single letter.
func isLetter(input string) bool {
	return len(input) == 1 && ('a' <= input[0] && input[0] <= 'z' || 'A' <= input[0] && input[0] <= 'Z')
}

// displayHangman displays a simple ASCII art representation of the hangman's progress.
func displayHangman(attempts, maxAttempts int) {
	switch attempts {
	case 1:
		fmt.Println("  ___ ")
		fmt.Println(" |   |")
		fmt.Println(" |   O")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println("_|_")
		
	}
}
