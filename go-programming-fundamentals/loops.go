//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func isDivisibleBy3(num int) bool {
	return num%3 == 0
}

func isDivisibleBy5(num int) bool {
	return num%5 == 0
}

func main() {
	for i := 1; i <= 50; i++ {
		switch num := i; {
		case isDivisibleBy3(num) && isDivisibleBy5(num):
			fmt.Println("FizzBuzz")
		case isDivisibleBy3(num):
			fmt.Println("Fizz")
		case isDivisibleBy5(num):
			fmt.Println("Buzz")
		default:
			fmt.Println(num)
		}
	}
}
