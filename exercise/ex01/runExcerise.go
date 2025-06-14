package ex01

import (
	"fmt"
	"math"
)

func RunExercise1() {
	fmt.Println("Hello, sleepy crocodile 🐊")

	fmt.Println("\nExercise 1: PrintName")
	PrintName("")
	PrintName("John Doe")

	fmt.Println("\nExercise 2: Print name and age")
	PrintNameAndAge("Jane Doe", 30)

	fmt.Println("\nExercise 3: Print numbers from 1 to 10")
	PrintNumbersUsingLoop(10, false)

	fmt.Println("\nExercise 4: Math power function")
	res := math.Pow(2, 11)
	fmt.Printf("2 raised to the power of 11 is: %.0f\n", res)

	fmt.Println("\nExercise 5: Print numbers using loop with new line")
	PrintNumbersUsingLoop(10, true)

	fmt.Println("\nExercise 6: Generate random number in range [1, 10]")
	fmt.Println(GenerateRandomInRange(1, 10))

	fmt.Println("\nExercise 7: Print current date")
	fmt.Println(PrintCurrentDate())

	var smallInt int8 = 100
	fmt.Println(smallInt + 29)
}
