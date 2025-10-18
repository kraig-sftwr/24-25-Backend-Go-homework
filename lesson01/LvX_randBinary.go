package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Binary(arr [100]int, tar int) int {
	L, R := 0, len(arr)-1

	for L <= R {
		M := (L + R) / 2

		if arr[M] == tar {
			return M + 1
		} else if arr[M] < tar {
			L = M + 1
		} else {
			R = M - 1
		}
	}

	return -1
}

func main() {

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)

	arr := [100]int{}
	for i := range arr {
		arr[i] = i + 1
	}

	y := Binary(arr, x)
	fmt.Println(x)
	fmt.Println(y)
}
