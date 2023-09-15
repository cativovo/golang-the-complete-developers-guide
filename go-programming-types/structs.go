//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	length int
	width  int
}

func calculateArea(rectangle Rectangle) int {
	return rectangle.length * rectangle.width
}

func calculatePerimeter(rectangle Rectangle) int {
	const Sides = 4
	return rectangle.length * Sides
}

func main() {
	rect1 := Rectangle{
		length: 4,
		width:  4,
	}

	fmt.Println("rect1 area", calculateArea(rect1))
	fmt.Println("rect1 perimeter", calculatePerimeter(rect1))

	rect2 := Rectangle{
		length: 8,
		width:  8,
	}

	fmt.Println("rect2 area", calculateArea(rect2))
	fmt.Println("rect2 perimeter", calculatePerimeter(rect2))
}
