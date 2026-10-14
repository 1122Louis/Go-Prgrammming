package main

import (
	"fmt"
	"strings"
)

// bill represents a single bill with named line items and a tip.
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// newBill creates a new bill with the given name and no items.
func newBill(name string) bill {
	return bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
}

// addItem adds a line item (name and price) to the bill.
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// setTip sets the tip amount on the bill.
func (b *bill) setTip(tip float64) {
	b.tip = tip
}

// total returns the sum of all items plus the tip.
func (b *bill) total() float64 {
	total := b.tip
	for _, price := range b.items {
		total += price
	}
	return total
}

// format returns a printable summary of the bill.
func (b *bill) format() string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "Bill breakdown for %s:\n", b.name)
	for name, price := range b.items {
		fmt.Fprintf(&sb, "%-20s ...$%.2f\n", name+":", price)
	}
	fmt.Fprintf(&sb, "%-20s ...$%.2f\n", "tip:", b.tip)
	fmt.Fprintf(&sb, "%-20s ...$%.2f", "total:", b.total())

	return sb.String()
}
