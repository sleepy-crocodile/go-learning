package ex02

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Problem1() {
	reader := bufio.NewReader(os.Stdin)
	for i := 1; i <= 3; i++ {
		line, err := reader.ReadString('\n') // Read input from the user
		if err != nil {
			panic(err)
		}
		fmt.Println(strings.TrimSpace(line))
	}

}
