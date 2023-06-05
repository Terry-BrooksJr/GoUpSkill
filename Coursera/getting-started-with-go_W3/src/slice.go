package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

func reAsk() bool {
	fmt.Println("Do You Have Any Integers to Add?  Y/N")
	fmt.Scan(&reAsk)
	if strings.ToLower(reAsk[0:1]) == "y" {
		return true 
	} else {
		return false
	}

}

func makeSlice(num ...int) []int {
	slice := make([]int, len(num))
	sort.Ints(slice)
	return slice
}

func collectNumbersFromUser() []int {
	var results []int = make([] int, 1)
	var userValue int
	var addNumbers bool
	addNumbers = reAsk()
	for (addNumbers) {
		addNumbers = reAsk()2
	}
	if addNumbers != true {
		
	}
	}
	fmt.Scan(&userValue)
	results = append(results, userValue)
return results

}

func main() {
	
}