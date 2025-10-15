package main

import "fmt"

func main() {
	var intArr [3]int32 // default value is [0,0,0]
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3]) // the execution is like python
	// value can also be changed by
	intArr[1] = 123
	fmt.Println(intArr[1])

	// Printing memory location - stored in contiguous memory location
	// The compiler doesn't store memory location of every element instead it derives it from first one, making it efficient
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// Pother ways of initialization
	var intArr1 [3]int32 = [3]int32{1,2,3}
	fmt.Println(intArr1)

	// Colon way
	intArr2 := [3]int32{1,2,3}
	intArr3 := [...]int32{1,2,3}
	fmt.Println(intArr2)
	fmt.Println(intArr3)
}

