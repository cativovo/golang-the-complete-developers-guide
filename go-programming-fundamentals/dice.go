//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these circumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
)

func rollDice(sides int) int {
	return rand.Intn(sides) + 1
}

func performDiceRoll(rolls, count, sides int) {

	for i := 1; i <= rolls; i++ {
		result := 0

		for j := 1; j <= count; j++ {
			rolled := rollDice(sides)
			result += rolled

			fmt.Println("Roll #", i, ", Die #", j, ":", rolled)
		}

		fmt.Println("Total rolled", result)

		switch result := result; {
		case result == 2 && count == 2:
			fmt.Println("Snake eyes")
		case result == 7:
			fmt.Println("Lucky 7")
		case isEven(result):
			fmt.Println("Even")
		case !isEven(result):
			fmt.Println("Odd")
		}

	}

}

func isEven(num int) bool {
	return num%2 == 0
}

func main() {
	const Rolls = 2
	const DiceCount = 2
	const DiceSide = 6

	performDiceRoll(Rolls, DiceCount, DiceSide)
}
