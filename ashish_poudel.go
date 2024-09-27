package main

import "fmt"

func agecount() {

	type Person struct {
		name string
		age  int
	}
	
	person1 := Person{"SHAN", 25}
	fmt.Println(person1)

	var person2 Person

	person2 = Person{
		name: "NIPUNA",
		age:  29,
	}

	fmt.Println(person2)
}
