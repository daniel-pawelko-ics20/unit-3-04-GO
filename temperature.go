// Copyright (c) 2021 Daniel Pawelko All rights reserved
//
// Created by: Daniel Pawelko
// Created on: May 2021
// This file contains GO program that converts Fahrenheit to celsius

package main

import (
	"fmt"
	"math"
)

// This main function will ask user for dimensions and output calculations
func main() {
	// Defining variables
	var fahren float64

	fmt.Println("Tempature: Fahrenheit to celsius")
	fmt.Println("(°F − 32) × 5/9 = °C")
	fmt.Println()

	// User Input
	fmt.Println("Please enter Fahrenheit:")
	fmt.Print("Fahrenheit: ")
	fmt.Scanln(&fahren)
	fmt.Print(fahren, "°F")
	fmt.Println()

	// calculations
	var celsius float64 = (fahren - 32) * 5 / 9

	// Print out volume
	fmt.Println("Celsius is: ", math.Round(celsius*100)/100)
}
