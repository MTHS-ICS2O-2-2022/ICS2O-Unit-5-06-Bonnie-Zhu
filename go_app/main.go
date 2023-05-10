// Created by: Bonnie Zhu
// Created on: May 2023
//
// This program multiplies two numbers together

package main

import (
	"fmt"
)

func main() {
	// This function multiplies two numbers together
	var number1 int
	var number2 int
	var counter = 0
	var answer = 0

	// input
	fmt.Println("Enter the first number: ")
	fmt.Scanln(&number1)
	fmt.Println("Enter the second number: ")
	fmt.Scanln(&number2)

	// process
	for counter < number2 {
		answer += number1
		counter++
	}

	// output
	fmt.Println("The answer is:", answer)

	fmt.Println("\nDone.")
}