package main

import (
	"fmt"
	"strings"
)



func stringSearch(input string) bool {
		var results bool
		if strings.ContainsAny(input,"ian"){
			results = true
			fmt.Println(results)
			return results
		} else {
			results = false
			fmt.Println(results)
			return results
		}	
}
func main( ) {
var userInput string
// var ianCheck bool	
fmt.Print("Please Enter a String: \n")
fmt.Scan(&userInput)
stringSearch(userInput)
main()

	
}