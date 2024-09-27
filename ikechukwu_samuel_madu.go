// package main

// import "fmt"

// // Define a function that accepts two parameters name(a string)and age(an integer) and prints them in a formatted way
// func greetStudent(name string, age string) {
// 	fmt.Printf("Hello, %s! I am %s years old.\n", name, age)
// }

package main

import (
	"fmt"
	"math"
)

func WholeNumbers() {
	number := 585225.0
	sqrt := math.Sqrt(number)
	fmt.Printf("The square root of %.0f is %.2f\n", number, sqrt)
}
