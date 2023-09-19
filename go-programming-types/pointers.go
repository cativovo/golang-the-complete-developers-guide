//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

type Item struct {
	name              string
	securityTagActive bool
}

func toggleSecurityTagStatus(prevStatus *bool, newStatus bool) {
	*prevStatus = newStatus
}

func checkout(items []Item) {
	for i := 0; i < len(items); i++ {
		toggleSecurityTagStatus(&items[i].securityTagActive, Inactive)
	}
}

func main() {
	items := []Item{
		{name: "item 1", securityTagActive: Active},
		{name: "item 2", securityTagActive: Active},
		{name: "item 3", securityTagActive: Active},
		{name: "item 4", securityTagActive: Active},
	}

	toggleSecurityTagStatus(&items[0].securityTagActive, Inactive)
	fmt.Println("Items:", items)

	checkout(items)
	fmt.Println("Items:", items)
}
