package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tips  float64
}

// add an item to the bill using receiver function
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// add tip to the bill using receiver function
func (b *bill) updateTip(tip float64) {
	b.tips = tip
}

func (b *bill) format() string {
	fs := "Bill breakdown \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		// fmt.Println(k, v)
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}
	// add tip
	fs += fmt.Sprintf("%25v ...%v \n", "tips:", b.tips)
	total += b.tips
	//total
	fs += fmt.Sprintf("%25v ...$%0.2f", "total:", total)
	return fs
}

func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("saveFile/bills/"+b.name+".txt", data, 0644) //returns an error if there is one, 0644 is the permission
	if err != nil {
		panic(err)
	}
	fmt.Println("The bill was saved to the file")
}
