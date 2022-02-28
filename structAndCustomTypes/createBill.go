package main

// create bill file is standalone
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type bill struct {
	name  string
	items map[string]float64
	tips  float64
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
	options, _ := getInput("Choose option (a -  add item,s - save bill,t - add tip):", reader)
	fmt.Println(options)
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
	fmt.Println(myBill)
}
