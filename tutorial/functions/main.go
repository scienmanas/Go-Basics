package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)
	// if err != nil {
	// 	fmt.Printf("The error encountered is: %v", err)
	// } else if remainder == 0 {
	// 	fmt.Printf("The result of the integer division is: %v", result)
	// } else {
	// 	fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	// }
	// using switch case
	switch {
	case err != nil:
		fmt.Printf("The error encountered is: %v", err)
	case remainder == 0:
		fmt.Printf("The result of the integer division is: %v", result)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v \n", result, remainder)
	}
	// One more way
	switch remainder {
	case 0:
		fmt.Println("The division was exact")
	case 1, 2:
		fmt.Println("the division was close")
	default:
		fmt.Println("The division was not close")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

// func intDivision(numerator int, denominator int) int {
// 	var result int = numerator/denominator
// 	return  result
// }

// To return multiple value
func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 9 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
