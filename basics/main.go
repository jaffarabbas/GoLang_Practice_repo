package main

import (
	"fmt"
	"go-basics/interfaces"
)

func printShapes(shape interfaces.Shape) {
	fmt.Println("Area: ", shape.Area())
}

func main() {
	fmt.Printf("Basics of go\n")
	testp := interfaces.Circle{Radius: 5}
	printShapes(testp)

	mt := interfaces.MyTest{}

	// Call the getTest method on the MyTest instance
	result := mt.getTest()

	// Print the result
	fmt.Println(result) // Output: test
}
