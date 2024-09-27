package main

import "fmt"

func agecount() {

	// declare a struct
	type Person struct {
		name string
		age  int
	}

	// assign value to struct while creating an instance
	person1 := Person{"SHAN", 25}
	fmt.Println(person1)

	// define an instance
	var person2 Person

	// assign value to struct variables
	person2 = Person{
		name: "NIPUNA",
		age:  29,
	}

	fmt.Println(person2)
}
