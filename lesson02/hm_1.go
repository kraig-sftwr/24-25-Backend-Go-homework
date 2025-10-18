package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func arrToMap(arr []int) map[int]int {
	m := make(map[int]int)

	for _, ele := range arr {
		m[ele]++
	}
	return m
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	l := s.Text()
	ps := strings.Fields(l)

	as := []int{}

	for _, p := range ps {
		a, _ := strconv.Atoi(p)
		as = append(as, a)
	}

	mp := arrToMap(as)

	for ele, tms := range mp {
		fmt.Printf("element:%d times:%d\n", ele, tms)
	}

}
