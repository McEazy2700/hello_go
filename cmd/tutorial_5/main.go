package main

import "fmt"

func main() {
	var p *uint32 = new(uint32) // define and initialize a pointer
	fmt.Printf("The pointer to address %v holds %v \n", p, *p)

	myArr := make([]uint32, 3, 20)
	add10Items(&myArr)
	fmt.Println(myArr)
}

func add10Items(arr *[]uint32) {
	for i := 0; i < 10; i++ {
		*arr = append(*arr, uint32(i))
	}
}
