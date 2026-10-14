package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// reader reads input from the terminal for the whole program.
var reader = bufio.NewReader(os.Stdin)

// getInput prints a prompt and returns the user's trimmed response.
func getInput(prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

// createBill asks for a name and returns a new, empty bill.
func createBill() bill {
	name, err := getInput("Create a new bill name: ")
	if err != nil {
		fmt.Println("Could not read name, using 'untitled'.")
		name = "untitled"
	}
	b := newBill(name)
	fmt.Println("Created the bill -", b.name)
	return b
}

// promptOptions drives the interactive menu until the user saves/quits.
func promptOptions(b bill) {
	for {
		opt, err := getInput("Choose option (a - add item, t - add tip, s - save/show): ")
		if err != nil {
			fmt.Println("Could not read option, stopping.")
			return
		}

		switch opt {
		case "a":
			addItemFlow(&b)
		case "t":
			addTipFlow(&b)
		case "s":
			fmt.Println()
			fmt.Println(b.format())
			return
		default:
			fmt.Println("That is not a valid option.")
		}
	}
}

func addItemFlow(b *bill) {
	name, err := getInput("Item name: ")
	if err != nil {
		fmt.Println("Could not read item name.")
		return
	}

	priceStr, err := getInput("Item price: ")
	if err != nil {
		fmt.Println("Could not read item price.")
		return
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		fmt.Printf("%q is not a valid price.\n", priceStr)
		return
	}

	b.addItem(name, price)
	fmt.Printf("Added %s ($%.2f) to the bill.\n", name, price)
}

func addTipFlow(b *bill) {
	tipStr, err := getInput("Tip amount: ")
	if err != nil {
		fmt.Println("Could not read tip.")
		return
	}

	tip, err := strconv.ParseFloat(tipStr, 64)
	if err != nil {
		fmt.Printf("%q is not a valid tip.\n", tipStr)
		return
	}

	b.setTip(tip)
	fmt.Printf("Set tip to $%.2f.\n", tip)
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
}
