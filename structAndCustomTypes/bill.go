package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tips  float64
}

func newBill(name string, tip float64) bill {
	b := bill{
		name: name,
		items: map[string]float64{
			"pie":  2.33,
			"cake": 4.55,
		},
		tips: tip,
	}

	return b
}

func (b bill) format() string {
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

// receiver function with pointers

// update the tip
func (b *bill) updateTip(tip float64) {
	b.tips = tip // we don't have to dereference, go will automatically dereference it
}

// add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
