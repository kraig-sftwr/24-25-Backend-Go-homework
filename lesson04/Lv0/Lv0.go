package main

import (
	"fmt"
	"time"
)

func main() {
	//time.Unix
	a := time.Now()
	b := a.Unix()
	fmt.Println(b)

	//when does defer conduct, before return or after?
	c := example()
	fmt.Println(c)

}

func example() int {
	defer fmt.Println("defer conducting..")
	fmt.Println("hello golang")
	return 100
}
