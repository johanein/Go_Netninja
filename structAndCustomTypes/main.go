package main

import (
	"fmt"
)

func main() {
	myBill := newBill("Albert", 6.22)
	myBill.updateTip(20)
	myBill.addItem("Mutton soup", 110.99)
	myBill.addItem("apple", 10.99)
	fmt.Println(myBill.format())
}
