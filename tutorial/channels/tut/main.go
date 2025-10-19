package main

import "fmt"

// func main() {
// 	var c = make(chan int)
// 	c <- 1
// 	var i = <-c
// 	// Why we get a deadlock error?
// 	// Because we are trying to send a value to a channel that is not ready to receive it.
// 	// We need to receive the value from the channel before sending it.
// 	// We can fix this by using a buffered channel.

// 	// Because when we write to an unbuffered channel, the code will be blocked until something else read from it so it will wait for forever unable to reach the line where we actually read from the channel.
// 	var c2 = make(chan int, 1)
// 	c2 <- 2
// 	var i2 = <-c2
// 	fmt.Println(i2)
// 	fmt.Println(i)
// }

func main() {
	var c = make(chan int)
	// var c = make(chan int, 5)
	go process(c)
	for i:= range c {
		fmt.Println(i)
	}
}

func process(c chan int) {
	defer close(c) // It will be closed when the function is finished
	for i := 0; i < 10; i++ {
		c <- i
	}
	// close(c)
}