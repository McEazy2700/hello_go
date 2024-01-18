package main

import "fmt"

func main() {
	intSlice := []int{1, 2, 40, 3}
	sum := sumSlices[int](intSlice)
	fmt.Println(sum)
	gasCar := Car[GasEngine]{
		make:  "Some make",
		model: "v20",
		engine: GasEngine{
			gallons: 20,
			mpg:     10,
		},
	}

	electricCar := Car[ElectricEngine]{
		make:  "Honda",
		model: "2022",
		engine: ElectricEngine{
			mpkwh: 300,
			kwh:   21,
		},
	}

	fmt.Println(gasCar, electricCar)
}

func sumSlices[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

type GasEngine struct {
	gallons uint32
	mpg     uint32
}

type ElectricEngine struct {
	mpkwh uint32
	kwh   uint32
}

type Car[T GasEngine | ElectricEngine] struct {
	make   string
	model  string
	engine T
}
