package main

import (
	"fmt"
	"strconv"
)

func addint() {
	var num1, num2 int64

	fmt.Print("Enter first decimal number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second decimal number: ")
	fmt.Scan(&num2)

	hex1 := strconv.FormatInt(num1, 16)
	hex2 := strconv.FormatInt(num2, 16)

	sumDecimal := num1 + num2
	sumHex := strconv.FormatInt(sumDecimal, 16)

	fmt.Printf("First number in hexadecimal: %s\n", hex1)
	fmt.Printf("Second number in hexadecimal: %s\n", hex2)
	fmt.Printf("Sum in decimal: %d\n", sumDecimal)
	fmt.Printf("Sum in hexadecimal: %s\n", sumHex)
}
