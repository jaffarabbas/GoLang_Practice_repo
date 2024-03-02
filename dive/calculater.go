package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func inputTaker() int64 {
	fmt.Println("Enter Number : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	selecter, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	return selecter
}

func sum(x, y int64) int64 {
	fmt.Print("Answer : ")
	return x + y
}
func sub(x, y int64) int64 {
	fmt.Print("Answer : ")
	return x - y
}
func divi(x, y int64) int64 {
	fmt.Print("Answer : ")
	return x / y
}
func multiply(x, y int64) int64 {
	fmt.Print("Answer : ")
	return x * y
}
func caluculater() {
	fmt.Println("*******CALCULATER*******\n")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter between + , - , / , * ")
	scanner.Scan()
	selecter := scanner.Text()
	if selecter == "+" {
		fmt.Println(sum(inputTaker(), inputTaker()))
	} else if selecter == "-" {
		fmt.Println(sub(inputTaker(), inputTaker()))
	} else if selecter == "/" {
		fmt.Println(divi(inputTaker(), inputTaker()))
	} else if selecter == "*" {
		fmt.Println(multiply(inputTaker(), inputTaker()))
	}
}
