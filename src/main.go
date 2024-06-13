// Package main implements a simple number guessing game in Go.
// The game generates a random number between 1 and 100, and the player is prompted to guess the number.
// The player receives feedback whether their guess is too small, too big, or correct. The game continues
// until the player guesses the correct number.
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// main is the entry point of the program. It starts the guessing game by calling the play function.
func main() {
	play()
}

// play starts the guessing game.
// It generates a random number between 1 and 100, then prompts the player to guess the number.
// The player receives feedback whether their guess is too small, too big, or correct.
// The game continues until the player guesses the correct number.
func play() {
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("Guess a number between 1 and 100")

	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(100) + 1

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Please input your guess: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid number!")
			continue
		}

		fmt.Printf("You guessed: %d\n", guess)

		if guess < secretNumber {
			fmt.Println("Too small!")
		} else if guess > secretNumber {
			fmt.Println("Too big!")
		} else {
			fmt.Println("You win!")
			break
		}
	}
}
