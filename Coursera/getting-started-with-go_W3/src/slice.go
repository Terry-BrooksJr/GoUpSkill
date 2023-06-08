package main

import (
	"fmt"
	"sort"
)

func main() {
	var sliceLength int
	fmt.Println("How Long is this Slice? 3 or 4 Digits")
	fmt.Scan(&sliceLength)
	if sliceLength == 4 {
		var num1, num2, num3, num4 int
		fmt.Println("What is the First Number?")
		fmt.Scan(&num1)
		fmt.Println("What is the Second Number?")
		fmt.Scan(&num2)
		fmt.Println("What is the Third Number?")
		fmt.Scan(&num3)
		fmt.Println("What is the Fourth Number?")
		fmt.Scan(&num4)
		result := []int{num1, num2, num3, num4}
		sort.Ints(result)
		fmt.Println(result)
	} else if sliceLength == 3 {
		var num1, num2, num3 int
		fmt.Println("What is the First Number?")
		fmt.Scan(&num1)
		fmt.Println("What is the Second Number?")
		fmt.Scan(&num2)
		fmt.Println("What is the Third Number?")
		fmt.Scan(&num3)
		result := []int{num1, num2, num3}
		sort.Ints(result)
		fmt.Println(result)
	}
	// trunk-ignore(golangci-lint/staticcheck)
	main()
}
