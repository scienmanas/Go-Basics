package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	var indexed = myString[0]
	fmt.Println(myString)
	fmt.Println(indexed)
	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	// Second option jsut use runes
	// var myString = []rune("résumé")
	// var indexed = myString[0]
	fmt.Println(myString)
	fmt.Println(indexed)
	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	var myRune = 'a'
	fmt.Printf("\nMy rune = %v", myRune)

	// String are immutable - we are creating new strings everytime
	// We an array is allocated internally and are appended when we called the writeString method and only at the end the string is made - much faster
	var strSlice = []string{"s", "u", "b", "s", "c"}
	// var catStr = ""
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v\n", catStr)

	// When using string builder the values are in array and then the string is made - much faster
}