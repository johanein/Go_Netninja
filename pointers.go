package main

import "fmt"

func updateNameUsingMemory(n *string) {
	*n = "johanein"
}

func main() {
	name := "Albert"
	// updateNameUsingMemory(name)

	fmt.Println("The memory address of name is", &name)
	m := &name
	fmt.Println("m is", m)
	fmt.Printf("value at memory in m is %q \n", *m)

	// passing pointer into the function
	updateNameUsingMemory(m)
	fmt.Println(name)

}
