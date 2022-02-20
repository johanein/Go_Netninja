package main

import "fmt"

// In a map all of the keys must be the same type and all of the values must be the same type
func main() {
	menu := map[string]float64{
		"soup":         8.532,
		"pie":          9.13,
		"salad":        6.98,
		"rice pudding": 4.453,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])
	for key, value := range menu {
		fmt.Printf("The key is \"%s\" and value is \"%v\" \n", key, value) // used ecsape characters here ;)
	}

	// maps with int as keys

	phoneBook := map[int]string{
		940053: "Harry potter",
	}
	fmt.Println(phoneBook[940053])

	phoneBook[940053] = "Albert" // updating value in a map

	for k, v := range phoneBook {
		fmt.Printf("Phonebook keys are '%v' and value '%s' \n", k, v)
	}
}
