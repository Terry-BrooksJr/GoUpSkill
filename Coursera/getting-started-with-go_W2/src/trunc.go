package main

import (
	"fmt"
)

func main() {
	// Global I/O Vars
	var userInput float64
	// Initial Prompt and Scan of Input
	fmt.Print("Please Provide a Number!:\n")
	fmt.Scan(&userInput)

	//Core Logic - Type Switch for the Input
	result := int64(userInput)
	if result == 0 && userInput != 0 {
		fmt.Printf("We had a Problem converting your entry into a 64-bit Int, Did you enter a Float Data Type?\n")
		main()
	} else {
		fmt.Printf("Here is Your Entry as a 64-Bit Int: %d \n", result)
		
		main()
	}
}
