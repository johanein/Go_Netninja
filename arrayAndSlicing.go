package main

import "fmt"

func main() {
	var array [3]int = [3]int{1, 2, 3} // original way
	// var array = [3]int{1, 2, 3} // shorthand way
	names := [4]string{"Arun", "Angel", "Albert", "Alfred"}
	fmt.Println(array, "length :", len(array))
	names[2] = "johanein"
	fmt.Println(names)

	// slices (length can be changed)
	scores := []int{9, 8, 7}
	scores = append(scores, 6)
	fmt.Println("scores :", scores)

	// slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[1:]
	rangeThree := names[:2]
	fmt.Println("ranges :", rangeOne)
	fmt.Println("rangeTwo :", rangeTwo)
	fmt.Println("rangeThree :", rangeThree)
}
