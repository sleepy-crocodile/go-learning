package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadLine() {
	reader := bufio.NewReader(os.Stdin)

	// Read a line(With space) from standard input
	delimeter := '\n'

	fmt.Print("Enter a line: ")
	line, err := reader.ReadString(byte(delimeter))
	line = strings.TrimSpace(line) // Trim whitespace from the line
	if err != nil {
		panic(err)
	}

	// Print the line read from standard input
	println("Line 1:", line)

	fmt.Print("Enter a line: ")
	line, err = reader.ReadString(byte(delimeter))
	line = strings.TrimSpace(line) // Trim whitespace from the line
	if err != nil {
		panic(err)
	}

	// Print the line read from standard input
	println("Line 2:", line)

}
