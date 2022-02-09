package main

import "fmt"

func main() {
	name := "Albert"
	age := 25
	// Print
	fmt.Print("Hello ")
	fmt.Print("again \n")
	// println
	fmt.Println("new line")
	fmt.Println("age is", age, "name is", name)
	// printf (formatted string) %_ = format specifiers
	fmt.Printf("My name is %v, and my age is %v \n", name, age)
	fmt.Printf("My name is %q, and my age is %v \n", name, age)
	fmt.Printf("My name is of type %T \n", age)
	fmt.Printf("Decimal points %f \n", 22.22)
	fmt.Printf("Decimal points limited %0.1f \n", 22.22)
	// Sprintf (save formatted strings)
	str := fmt.Sprintf("My name is %v, and my age is %v \n", name, age)
	fmt.Println("saved string is :", str)
}
