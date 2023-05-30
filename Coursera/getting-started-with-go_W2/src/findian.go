package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
var userInput string
func stringSearch(userInput string) string {
	fmt.Print("Enter the test string: ")
	reader := bufio.NewReader(os.Stdin)
		userInput, _ = reader.ReadString('\n')
		userInput = strings.ToLower(strings.TrimSpace(userInput))
		if strings.HasPrefix(userInput, "i") &&
			strings.Contains(userInput, "a") &&
			strings.HasSuffix(userInput, "n") {
			return ("Found!")
		} else {
			return ("Not Found!")
		}
}

func main() {
	fmt.Println(stringSearch(userInput))
	main()

}
