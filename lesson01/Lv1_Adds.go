package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	var a int
	var b int

	fmt.Scanln(&a, &b)

	fmt.Println(add(a, b))

}
