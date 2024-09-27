package main

import "fmt"

func main() {
	fmt.Println("Welcome to Group A's Week 4 Project!")
	print("Nipuna Karunarathna \n input needed ")
	var url string
	fmt.Print("Enter URL: ")
	fmt.Scanln(&url)
	status := CheckHostAvailability(url)
	fmt.Print(status)
	print("Seruban Peter Shan \n")

	SerubanPeterShan(8)

	print("Ashish Poudel \n")

	addint()

	// greetStudent("Sam", "40")

	WholeNumbers()

}
