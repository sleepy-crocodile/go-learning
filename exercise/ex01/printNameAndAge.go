package ex01

import "fmt"

func PrintNameAndAge(name string, age int) {
	if name == "" {
		name = "Sleepy Crocodile 🐊"
	}
	if age <= 0 {
		age = 1 // Default age if not provided
	}
	// Print the name and age
	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
}
