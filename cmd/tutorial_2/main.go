package main

import (
	"errors"
	"fmt"
)

func main() {
	printThis("something")
	val, rem, err := intDivision(20, 12)
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case rem > 0:
		fmt.Printf("We have a remainder of %v", rem)
	default:
		fmt.Printf("The value of the divisio is %v and the remainder is %v", val, rem)
	}
}

func printThis(val string) {
	fmt.Println(val)
}

func intDivision(num int, den int) (int, int, error) {
	var err error
	if den == 0 {
		err = errors.New("Cannont divide by zero")
		return 0, 0, err
	}
	divided := num / den
	remainder := num % den
	return divided, remainder, err
}
