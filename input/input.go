package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadNumber() {
	reader := bufio.NewReader(os.Stdin)

	// Read a line(With space) from standard input
	delimeter := '\n'

	fmt.Print("Enter a line: ")
	line, err := reader.ReadString(byte(delimeter))
	line = strings.TrimSpace(line) // Trim whitespace from the line
	if err != nil {
		panic(err)
	}

	// For strinToNumber.txt
	num, numerr := strconv.Atoi(line)
	if numerr != nil {
		panic(numerr)
	}
	// Print the line read from standard input
	println("Line 1:", line, "Integer: ", num)

	fmt.Print("Enter a line: ")
	line, err = reader.ReadString(byte(delimeter))
	line = strings.TrimSpace(line) // Trim whitespace from the line
	if err != nil {
		panic(err)
	}

	// For strinToNumber.txt
	num, numerr = strconv.Atoi(line)
	if numerr != nil {
		panic(numerr)
	}

	// Print the line read from standard input
	println("Line 2:", line, "Integer: ", num)

	fmt.Print("Enter a line: ")
	line, err = reader.ReadString(byte(delimeter))
	line = strings.TrimSpace(line) // Trim whitespace from the line
	if err != nil {
		panic(err)
	}

	// For strinToNumber.txt
	floatNum, floatErr := strconv.ParseFloat(line, 64)
	if floatErr != nil {
		panic(floatErr)
	}

	// Print the line read from standard input
	println("Line 2:", line, "Integer: ", floatNum)

}
