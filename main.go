package main

import (
	"fmt"
	"time"
)

func main() {
	// Load IST timezone
	location, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	// Get current time in IST
	currentTime := time.Now().In(location)

	// Print time
	fmt.Println("Current time in IST:", currentTime.Format("2006-01-02 15:04:05 MST"))
}
