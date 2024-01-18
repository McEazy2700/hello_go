package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2) // Integer division results in an integer and the result is rounded down
	fmt.Println(intNum1 % intNum2) // To get the remainder use a percent or modulo sign

	var myString string = "Hello World"           // strings can be created with double quotes '"' or backticks "`"
	fmt.Println(myString)                         // you can add a new line to quote strings with \n backtick strings are multiline
	fmt.Println(len(myString))                    // using len to count strings returns the amount of bytes
	fmt.Println(utf8.RuneCountInString(myString)) // use this to get the number of characters (runes) in go

	var myRune rune = 'a' // A rune represents a character
	fmt.Println(myRune)

	var myBoolean bool = true // classic
	fmt.Println(myBoolean)

	// Go sets default values for undeclared variables like intNum
	// the default value depends on its type
	// for all number like types is 0, for strings its an empty string '' and for booleans its false
	var myString2 = "foo" // you can omit specifying the type if you set the value right away (inference)
	fmt.Println(myString2)
	myNum := 22 // You can also omit the var if you use the := short hand (kind of looks suspicious)
	fmt.Println(myNum)

	var num1, num2 int = 23, 24 // initialize multiple variables at once (Nice)
	fmt.Println(num1, num2)
	// Add the type when it isn't obvious `var myVar int := foo()`
	const hoursInADay = 24 // constants never change (standard behavious)
	fmt.Println(hoursInADay)
}
