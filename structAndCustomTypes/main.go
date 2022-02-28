package main

import (
	"fmt"
)

func main() {
	myBill := newBill("Albert")
	myBill.updateTip(5)
	myBill.addItem("Mutton soup", 20.99)
	myBill.addItem("apple", 2.99)
	fmt.Println(myBill.format())
}
