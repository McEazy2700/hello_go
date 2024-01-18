package main

import "fmt"

type GasEngine struct {
	mpg     uint8
	gallons uint8
}

type ElectricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e ElectricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func (e GasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) bool {
	return miles <= e.milesLeft()
}

func main() {
	var myEngine GasEngine = GasEngine{mpg: 20, gallons: 30}
	fmt.Println(myEngine.mpg, myEngine.gallons)
	engines := map[string]GasEngine{"John": {20, 30}}
	engines["Sally"] = GasEngine{21, 21}
	fmt.Println(engines)
	for _, val := range engines {
		fmt.Println(val.milesLeft())
		fmt.Printf("Cam make it a 250 mile is %v \n", canMakeIt(val, uint8(250)))
	}
}
