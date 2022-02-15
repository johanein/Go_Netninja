package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	var upperString = strings.ToUpper(n)
	var splitString = strings.Split(upperString, " ")

	var initials []string
	for _, v := range splitString {
		initials = append(initials, v[:1]) // I dont know why we can't use v[0] here. One day when you know save it here.
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func main() {
	first, second := getInitials("albert c y")
	fmt.Printf("first: %s, second: %s \n", first, second)
	var first1, second1 = getInitials("roopesh")
	fmt.Printf("first1: %s, second1: %s \n", first1, second1)
}
