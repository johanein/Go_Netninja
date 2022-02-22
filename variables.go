package main // the package name we created

import "fmt" // package from go standard library

func main() { // entry point of our application
	var nameOne string = "Albert"
	var nameTwo = "Lify" // here it infers the type
	var nameThree string // value given at a later point of time

	nameFour := "Banner" // shorthand version, this can only be used inside of a function

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)
	nameOne = "Apple"
	nameTwo = "Orange"
	nameThree = "Grapes"
	nameFour = "Lotus"
	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	// numbers
	var ageOne int = 30
	ageTwo := 40
	var numOne int8 = 127
	var numTwo int8 = -128
	var numThree uint8 = 255
	var scoreOne float32 = 23.3
	var scoreTwo float64 = 23345575675647357535.3 // 64 commonly used
	scoreThree := 1.2

	fmt.Println(ageOne, ageTwo, numOne, numTwo, numThree, scoreOne, scoreTwo, scoreThree)

}
