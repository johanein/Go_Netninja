package main

import "fmt"

func updateName(n string) string {
	n = "johanein"
	return n
}

func updateMenu(m map[string]float64) {
	m["beef"] = 1000.99
	m["beer"] = 666.66
}

func main() {
	// group A types -> strings, ints, bools, floats, array, structs
	name := "Albert"

	name = updateName(name)
	fmt.Println(name)

	// group B types -> slices, maps, function (Pointer wrapper values)
	menu := map[string]float64{
		"egg":    3.44,
		"carrot": 5.66,
		"beef":   6.22,
	}
	updateMenu(menu)
	for k, v := range menu {
		fmt.Println("key: ", k, "value: ", v)
	}

}
