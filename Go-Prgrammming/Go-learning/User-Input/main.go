package main

import "fmt"

func main() {
	mybill := newBill("Louis Soup")

	mybill.addItem("onion soup", 5.99)
	
	fmt.Println(mybill)