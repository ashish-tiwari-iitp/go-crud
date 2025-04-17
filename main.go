package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printTimes() {
	// Load IST and GMT timezones
	istLocation, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println("Error loading IST location:", err)
		return
	}

	gmtLocation, err := time.LoadLocation("GMT")
	if err != nil {
		fmt.Println("Error loading GMT location:", err)
		return
	}

	// Get current time
	now := time.Now()

	// Print times
	fmt.Println("Current time in UTC:", now.UTC().Format("2006-01-02 15:04:05 MST"))
	fmt.Println("Current time in IST:", now.In(istLocation).Format("2006-01-02 15:04:05 MST"))
	fmt.Println("Current time in GMT:", now.In(gmtLocation).Format("2006-01-02 15:04:05 MST"))
}

func playGuessingGame() {
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(10) + 1 // number between 1 and 10
	var guess int

	fmt.Println("\n🎮 Let's play a guessing game!")
	fmt.Println("I'm thinking of a number between 1 and 10. Can you guess it?")

	for attempts := 0; attempts < 3; attempts++ {
		fmt.Printf("Attempt %d: Enter your guess: ", attempts+1)
		fmt.Scan(&guess)

		if guess == secretNumber {
			fmt.Println("🎉 Correct! You win!")
			return
		} else if guess < secretNumber {
			fmt.Println("Too low!")
		} else {
			fmt.Println("Too high!")
		}
	}

	fmt.Printf("😢 Sorry, you lost. The number was %d.\n", secretNumber)
}

func main() {
	printTimes()
	playGuessingGame()
}
