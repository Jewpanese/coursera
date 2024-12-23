package main

import (
	"fmt"
	"math"
)

func main() {

	var floatingNumber float64
	fmt.Print("Please enter a floating number: \r\n")

	if _, err := fmt.Scan(&floatingNumber); err != nil {
		fmt.Println(err)
	}

	truncNumber := math.Trunc(floatingNumber)
	fmt.Println(truncNumber)
}
