package main

import (
	"fmt"
	"math"
)

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	ceil := int(math.Sqrt(float64(n)))

	for i := 3; i <= ceil; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var x int

	fmt.Scanln(&x)

	fmt.Println(IsPrime(x))

}
