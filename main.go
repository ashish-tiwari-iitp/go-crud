package main

import (
	"fmt"
	"time"
)

func main() {
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

	// Format and print time in UTC
	fmt.Println("Current time in UTC:", now.UTC().Format("2006-01-02 15:04:05 MST"))

	// Format and print time in IST
	fmt.Println("Current time in IST:", now.In(istLocation).Format("2006-01-02 15:04:05 MST"))

	// Format and print time in GMT
	fmt.Println("Current time in GMT:", now.In(gmtLocation).Format("2006-01-02 15:04:05 MST"))
}
