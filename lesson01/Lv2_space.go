package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	fmt.Scanln(&r)

	s := math.Pi * math.Pow(r, 2)

	fmt.Println(s)

}
