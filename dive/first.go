package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	scanner.Scan()
	input2, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println(input + input2)
}
