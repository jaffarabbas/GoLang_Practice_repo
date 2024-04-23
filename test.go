// Package declaration
package main

// Import packages
import (
	"fmt"
)

// Main function
func main() {
	num := 1
	fmt.Println(num)

	fmt.Println(dynamicVariables(1, 2, 3, 4, 5))
}

func test(a int, b string) (int, int) {
	if b == "1" {
		return 1, 2
	} else {
		return 3, 4
	}
}

func dynamicVariables(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
