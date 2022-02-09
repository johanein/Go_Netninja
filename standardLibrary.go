package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	stringVar := "Hello there, how are you ?"
	fmt.Println("Does the stringVar contains the searching values :", strings.Contains(stringVar, "Hello"))
	fmt.Println("Using replaceAll method :", strings.ReplaceAll(stringVar, "Hello", "Hi"))
	fmt.Println("Using ToUpper method :", strings.ToUpper(stringVar))
	fmt.Println("Find the index of ho :", strings.Index(stringVar, "ho"))
	fmt.Println("Splitting the stringVar :", strings.Split(stringVar, "e"))

	// The original value is unchanged
	fmt.Println(" The original stringVar value :", stringVar)

	ages := []int{60, 100, 70, 50, 40, 20, 30}
	sort.Ints(ages)
	index := sort.SearchInts(ages, 40)
	fmt.Println(" The searched index value :", index)
	fmt.Println(" The original array value :", ages)

	namesArray := []string{"Arun", "Angel", "Albert", "Alfred"}
	sort.Strings(namesArray)
	fmt.Println(sort.SearchStrings(namesArray, "Arun"))
	fmt.Println(namesArray)

}
