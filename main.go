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

var leaderboard = make([]string, 0)

func main() {
	rand.Seed(time.Now().UnixNano())
	var playAgain string

	fmt.Println("Welcome to the Ultimate Number Guessing Game!")
	for {
		// Display the leaderboard
		displayLeaderboard()

		// Select difficulty level
		level := selectLevel()

		// Adjust the range based on the difficulty level
		var maxRange int
		switch level {
		case 1:
			maxRange = 50
		case 2:
			maxRange = 100
		case 3:
			maxRange = 200
		}

		// Generate the secret number
		secretNumber := rand.Intn(maxRange) + 1

		// Start the game
		fmt.Printf("I'm thinking of a number between 1 and %d.\n", maxRange)
		fmt.Println("Can you guess what it is? Let's go!")

		// Track the number of attempts and start the timer
		var attempts int
		var startTime = time.Now()

		// Main game loop
		for {
			// Read user input
			fmt.Print("Enter your guess: ")
			input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			input = strings.TrimSpace(input)

			// Convert input to an integer
			guess, err := strconv.Atoi(input)
			if err != nil || guess < 1 || guess > maxRange {
				fmt.Println("Please enter a valid number between 1 and", maxRange)
				continue
			}

			attempts++

			// Check if the guess is correct
			if guess < secretNumber {
				fmt.Println("Too low! Try again.")
				provideHint(secretNumber, guess)
			} else if guess > secretNumber {
				fmt.Println("Too high! Try again.")
				provideHint(secretNumber, guess)
			} else {
				// User guessed correctly
				elapsedTime := time.Since(startTime)
				fmt.Printf("🎉 Congratulations! You guessed the number in %d attempts.\n", attempts)

				// Calculate score
				score := calculateScore(attempts, elapsedTime)
				fmt.Printf("Your score: %d\n", score)

				// Add score to leaderboard
				leaderboard = append(leaderboard, fmt.Sprintf("Score: %d (Time: %.2fs, Attempts: %d)", score, elapsedTime.Seconds(), attempts))

				break
			}
		}

		// Ask if the user wants to play again
		fmt.Print("\nDo you want to play again? (y/n): ")
		fmt.Scanln(&playAgain)

		if strings.ToLower(playAgain) != "y" {
			fmt.Println("Thanks for playing! See you next time.")
			break
		}
	}
}

// Function to display the leaderboard
func displayLeaderboard() {
	fmt.Println("\nLeaderboard:")
	if len(leaderboard) == 0 {
		fmt.Println("No scores yet. Play a game to add your name to the leaderboard!")
	} else {
		for i, score := range leaderboard {
			fmt.Printf("%d. %s\n", i+1, score)
		}
	}
	fmt.Println()
}

// Function to select difficulty level
func selectLevel() int {
	var level int
	for {
		fmt.Println("\nSelect difficulty level:")
		fmt.Println("1: Easy (1-50)")
		fmt.Println("2: Medium (1-100)")
		fmt.Println("3: Hard (1-200)")
		fmt.Print("Enter the level number (1/2/3): ")

		// Read input
		fmt.Scanln(&level)
		if level >= 1 && level <= 3 {
			break
		} else {
			fmt.Println("Invalid choice. Please select a valid level.")
		}
	}
	return level
}

// Function to calculate score based on attempts and time
func calculateScore(attempts int, elapsedTime time.Duration) int {
	// Basic score logic
	baseScore := 100 - (attempts-1)*10 - int(elapsedTime.Seconds())

	// Ensure score doesn't go negative
	if baseScore < 0 {
		baseScore = 0
	}

	// Apply bonus for fast guesses
	if elapsedTime.Seconds() < 10 {
		baseScore += 20
		fmt.Println("Bonus: You guessed really fast!")
	}

	// Apply bonus for few attempts
	if attempts <= 3 {
		baseScore += 30
		fmt.Println("Bonus: You guessed it quickly with few attempts!")
	}

	return baseScore
}

// Function to provide hints
func provideHint(secretNumber, guess int) {
	if guess < secretNumber {
		fmt.Println("Hint: The secret number is higher than your guess.")
	} else if guess > secretNumber {
		fmt.Println("Hint: The secret number is lower than your guess.")
	}
}