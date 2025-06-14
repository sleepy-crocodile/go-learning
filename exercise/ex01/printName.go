package ex01

import "fmt"

func PrintName(name string) {
	if name == "" {
		name = "Sleepy Crocodile 🐊"
	}
	// Print the name
	fmt.Println("Hello,", name)
}
