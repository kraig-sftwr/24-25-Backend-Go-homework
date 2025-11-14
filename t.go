package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MaxBubble(b int, slice []int) []int {

}

func main() {
	slice1 := []int{}

	s1in := bufio.NewScanner(os.Stdin)
	s1in.Scan()
	s1 := s1in.Text()

	s1ps := strings.Fields(s1)

	for _, s1p := range s1ps {
		num, _ := strconv.Atoi(s1p)
		slice1 = append(slice1, num)
	}

	slice2 := []int{}

	fmt.Println(slice2)

}
