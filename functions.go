package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Hello %s \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func returnGreeting(n string) string {
	return fmt.Sprintf("Returning %v", n)
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayGreeting("Albert")

	namesArray := []string{"Arun", "Angel", "Albert", "Alfred"}
	cycleNames(namesArray, sayGreeting)

	fmt.Println(returnGreeting("Albert"))
	fmt.Println("The circle Area is :", circleArea(10.5))
	fmt.Printf("The circle Area is %0.3f \n", circleArea(10.5))
}
