package main

import (
	"fmt"
	"time"
)

func main() {
	// Load IST timezone
	istLocation, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println("Error loading IST location:", err)
		return
	}

	// Get current time
	now := time.Now()

	// Format and print time in UTC
	fmt.Println("Current time in UTC:", now.UTC().Format("2006-01-02 15:04:05 MST"))

	// Format and print time in IST
	fmt.Println("Current time in IST:", now.In(istLocation).Format("2006-01-02 15:04:05 MST"))
}
