package main

import "fmt"
import "unicode/utf8"

func main() {
	var intNum int16 = 32767
	intNum = intNum + 1
	// var intNum uint16 = 65535
	// intNum = intNum + 1
	fmt.Println(intNum) // it is in negative bcoz we are exceeding the

	var floatNum float32 = 12433.3432
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2) // integer type returned
	fmt.Println(intNum1%intNum2)

	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)

	fmt.Println(len("test")) // 4 -> its number of bytes not number of characters in the string
	fmt.Println(utf8.RuneCountInString("y"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	// We can not use the datatype keyword and directy assign, it will understand on its own
	var myVar1 = "text"
	myVar2 := "text"
	fmt.Println(myVar1)
	fmt.Println(myVar2)

	// Multiple variables initialization
	// var var1, var2 int = 1,2
	var1, var2 := 1,2
	fmt.Println(var1)
	fmt.Println(var2)

	const myConst string = "constant value"
	fmt.Println(myConst)
}

