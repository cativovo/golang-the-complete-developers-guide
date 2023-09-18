//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printContents(parts []Part) {
	for i := 0; i < len(parts); i++ {
		part := parts[i]
		fmt.Println("Part:", part)
	}
}

func main() {
	assembyLine := []Part{"part 1", "part 2", "part 3"}
	fmt.Println("3 parts:")
	printContents(assembyLine)

	fmt.Println("\nAdded two parts:")
	assembyLine = append(assembyLine, "part 4", "part 5")
	printContents(assembyLine)

	fmt.Println("\nSliced:")
	assembyLine = assembyLine[3:]
	printContents(assembyLine)
}
