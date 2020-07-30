package main

import "fmt"

var a = 50
var b = 954

func main() {
	var Add int = a + b
	fmt.Printf("Addition of two numbers is: %d\n", Add)
	var Substract int = a - b
	fmt.Printf("Substraction of two numbers is: %d\n", Substract)
	var Multiply = a * b
	fmt.Printf("Multipicatiob of two numbers is:%d\n", Multiply)
	var div = b / a
	fmt.Printf("Division:%d\n", div)
	var remainder = b % a
	fmt.Printf("remainder is:2%d\n", remainder)
}
