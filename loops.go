package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("loops")
	y := 0
	// while loop in go
	for y < 5 {
		fmt.Println("value of y is :", y)
		y++
	}
	// for loop in go
	for x := 5; x < 10; x++ {
		fmt.Println("value of x is :", x)
	}

	namesArray := []string{"Arun", "Angel", "Albert", "Alfred"}
	sort.Strings(namesArray)
	for i := 0; i < len(namesArray); i++ {
		fmt.Println(namesArray[i])
	}

	// same as 'for in' in js
	for key, value := range namesArray {
		fmt.Printf("The index is %v, and the value is %s \n", key, value)
	}
	// If only valus is needed
	for _, value := range namesArray {
		fmt.Printf("The value is %s \n", value)
		value = "updated value"
	}

	fmt.Println(namesArray)
}
