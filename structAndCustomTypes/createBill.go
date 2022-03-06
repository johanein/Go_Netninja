package main

// create bill file is standalone
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tips:  0,
	}

	return b
}

func getInput(prompt string, r *bufio.Reader) (string, error) {

	fmt.Print(prompt)
	input, errorVar := r.ReadString('\n')
	return strings.TrimSpace(input), errorVar
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin) // Stdin is standard input which is the terminal
	name, _ := getInput("Enter the bill name:", reader)
	b := newBill(name)
	fmt.Println("created the bill with name :", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin) // Stdin is standard input which is the terminal
	options, _ := getInput("Choose option (a -  add item,s - save bill,t - add tip,e - escape from loop):", reader)
	switch options {
	case "a":
		name, _ := getInput("Item name:", reader)
		price, _ := getInput("Item Price:", reader)
		floatedString, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("price should be a number")
			promptOptions(b)
		}
		b.addItem(name, floatedString)
		promptOptions(b)
	case "s":
		fmt.Println("you chose to save the bill", b)
	case "t":
		tip, _ := getInput("Enter tip amount ($)", reader)
		floatedTip, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Tip should be a number")
			promptOptions(b)
		}
		b.updateTip(floatedTip)
		promptOptions(b)
	case "e":
	default:
		fmt.Println("invalid option....")
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}
