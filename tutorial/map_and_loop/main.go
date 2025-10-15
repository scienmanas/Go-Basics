package main

import (
	"fmt"
)

// map[string]int32
func main() {
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Sarah"])
	// fmt.Println(myMap2["Manas"]) // gives zero no error, it always return something
	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is: %v\n", age)
	} else {
		fmt.Println("Invalid name")
	}

	// Loops
	for name, age := range myMap2 {
		fmt.Printf("Name: %v & Age: %v\n", name, age)
	}

	// Normal printing
	// Shorthand like i-- and all is supported
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
}
