package main

import (
	"jalvarez/golang-algorithms/mypackage"
	"fmt"
)

func main() {
	var myCar mypackage.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2022
	fmt.Println(myCar)
}