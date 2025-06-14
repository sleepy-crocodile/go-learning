package ex01

import (
	"time"
)

/**
 * PrintName prints the given name.
 * time is the complex data type.
 * @param name string - The name to print.
 */
func PrintCurrentDate() string {
	return time.Now().String()
}
