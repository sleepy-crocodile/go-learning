package exercise

import (
	"fmt"
	"math/rand/v2"
)

func PrintNumbersUsingLoop(max int, newLine bool) {
	for i := 1; i <= max; i++ {
		if i == 1 {
			if newLine {
				fmt.Println(i)
			} else {
				fmt.Print(i)
			}
			continue
		}
		if newLine {
			fmt.Println(i)
		} else {
			fmt.Print(" ", i)
		}
	}
	if !newLine {
		fmt.Println() // Print a newline at the end
	}
}

func GenerateRandomInRange(min int, max int) int {
	return rand.IntN(max-min+1) + min
}
