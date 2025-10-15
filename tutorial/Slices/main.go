package main

import (
	"fmt"
	"time"
)

func main() {
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\nThe length is %v with capacity %v \n", len(intSlice), cap(intSlice))
	// The capacity is two more but we cannot access it.

	// Multiple append can also be done
	var intSlice2 []int32 = []int32{8, 9}
	intSlice2 = append(intSlice, intSlice2...) // using the spread operator
	fmt.Println(intSlice2)

	// Using make operator
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)
	println(intSlice3)

	sliceOperation()
}

func sliceOperation() {
	var n int = 1000000
	var testSlice = []int{}
	// var testSlice2 = make([]int, 0, n)
	testSlice2 := make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v", timeLoop(testSlice2, n))
}


func timeLoop (slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}