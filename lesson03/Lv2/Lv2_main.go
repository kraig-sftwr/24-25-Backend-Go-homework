package main

import (
	"Lv2/utils"
	"fmt"
)

func main() {
	fmt.Printf("Please input an string: ")
	var s1 string
	fmt.Scan(&s1)

	s2 := utils.Reverse(s1)
	res := utils.IsPalindrome(s1, s2)

	if res == true {
		fmt.Println("Yes it is a palindrome")
	} else {
		fmt.Println("No it isn't a palindrome")
	}
}
