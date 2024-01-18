package main

import "fmt"

func main() {
	var intArr [3]int32
	fmt.Println(intArr[1])
	fmt.Println(intArr[:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var initArr [4]int32 = [4]int32{2, 3, 5}
	fmt.Println(initArr)
	inferedArr := [...]int32{2, 3, 5}
	fmt.Println(inferedArr)

	var intSlice []int32 = []int32{2, 39, 4} // by omiting the length of the array we crate a slice
	fmt.Printf("The length of the slice is %v and its capacity is %v \n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 30)
	fmt.Printf("The length of the slice is %v and its capacity is %v \n", len(intSlice), cap(intSlice))

	otherSlice := []int32{23, 50}
	intSlice = append(intSlice, otherSlice...) // spreading (Nice!)
	fmt.Printf("The length of the slice is %v and its capacity is %v \n", len(intSlice), cap(intSlice))

	var madeSlice []int32 = make([]int32, 3, 8) // make a slice with an optional capacity
	fmt.Println(madeSlice)

	map1 := make(map[string]uint8) // key value pairs
	fmt.Println(map1)
	map1["john"] = 32
	fmt.Println(map1)
	initMap := map[string]uint8{"John": 20, "Sarah": 32}
	fmt.Println(initMap["John"])
	fmt.Println(initMap["Peter"])     // map always returns a value in this case the default value of our value type uint8
	value, exists := initMap["Jacob"] // use the second boolean to confirm if a value is in the map
	fmt.Printf("Exists status for record Jacob is %v and it's value is %v \n", exists, value)
	delete(initMap, "John") // removes an item from a map
	val, exists := initMap["John"]
	if exists {
		fmt.Printf("John's record still exists '%v'. Well this is akward", val)
	} else {
		fmt.Println("All records of John are gone")
	}

	for name, age := range initMap { // iterate over a map
		fmt.Printf("%v is %v years old \n", name, age)
	}

	for i, val := range initArr {
		fmt.Printf("%v is at index %v or address %v \n", val, i, &initArr[i])
	}

	// While
	for count := 0; count < 10; count++ { // Classic for loop
		intSlice = append(intSlice, int32(count+1))
	}
	fmt.Println(intSlice)
}
