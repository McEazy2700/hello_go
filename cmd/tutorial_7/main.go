package main

import "fmt"

func main() {
	c := make(chan int, 30)
	go process(c, 30, 30)
	for i := range c {
		fmt.Println(i)
	}
}

func process(c chan int, val int, amt int) {
	defer close(c)
	for i := 0; i < amt; i++ {
		c <- val * i
	}
	fmt.Println("Process complete")
}
