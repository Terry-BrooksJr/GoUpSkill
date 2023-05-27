package main
import (
	"strings"
	"fmt"
)

func main( ) {
	var input string
	var ianCheck Bool
	fmt.Println("Please Enter a String")
	fmt.Scan(&input)

	if strings.ContainsAny("ian"){
		ianCheck = true
		fmt.Print(ianCheck)
		return ianCheck
	} else {
		ianCheck = false
		fmt.Print(ianCheck)
		return ianCheck
	}
}