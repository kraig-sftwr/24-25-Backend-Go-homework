package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	//start timing1
	st1 := time.Now()

	//create file a(with buffer)
	a, _ := os.Create("a.txt")
	//create a NewWriter
	a_writer := bufio.NewWriter(a)
	for i := 0; i < 1000; i++ {
		a_writer.WriteString("Hello World!\n")
	}
	a_writer.Flush()
	a.Close()
	fmt.Println(time.Since(st1)) //end and output timing1

	//create file b(without buffer)
	st2 := time.Now()
	b, _ := os.Create("b.txt")
	for i := 0; i < 1000; i++ {
		b.WriteString("Hello World!\n")
	}
	b.Close()
	fmt.Println(time.Since(st2)) //end and output timing2

}
