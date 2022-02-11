package main

import "fmt"

func main() {
	age := 25
	println(age <= 50)
	println(age >= 50)
	println(age == 45)
	println(age != 50)

	// else if
	if age < 30 {
		fmt.Println("The person is a Youth")
	} else if age < 40 {
		fmt.Println("The person is in Middle age")
	} else {
		fmt.Println("The person is an Adult")
	}

	namesArray := []string{"Arun", "Angel", "Albert", "Alfred"}
	// using break and continue
	for index, value := range namesArray {
		if index == 1 {
			fmt.Printf("Continuing at pos %v with value %s \n", index, value)
			continue
		}
		if index == 2 {
			fmt.Printf("Breaking at pos %v with value %s \n", index, value)
			break
		}
		fmt.Printf("the value at pos %v is %s \n", index, value)
	}
}
