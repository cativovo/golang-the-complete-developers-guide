//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price float64
}

func printShoppingListInformation(list [4]Product) {
	cost := 0.0
	totalItems := 0

	for i := 0; i < len(list); i++ {
		product := list[i]
		cost += product.price

		if product.name != "" {
			totalItems += 1
		}
	}

	fmt.Println("Last item is", list[totalItems-1])
	fmt.Println("Total items", totalItems)
	fmt.Printf("total cost %.2f\n", cost)
}

func main() {
	shoppingList := [4]Product{}

	shoppingList[0] = Product{
		name:  "apple juice",
		price: 4.99,
	}

	shoppingList[1] = Product{
		name:  "bread",
		price: 6.00,
	}

	shoppingList[2] = Product{
		name:  "strawberry jam",
		price: 10.00,
	}

	printShoppingListInformation(shoppingList)

	shoppingList[3] = Product{
		name:  "nuts",
		price: 1.69,
	}

	printShoppingListInformation(shoppingList)
}
