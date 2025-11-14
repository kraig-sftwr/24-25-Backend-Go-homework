package main

import (
	"fmt"
	"sync"
)

func testgroup() {
	var wg sync.WaitGroup //go并发编程中的重要结构体WaitGroup

	wg.Add(5)

	go func() {
		fmt.Println("Hello goroutine111")
		wg.Done() //计数器-1
	}() //这里圆括号的意义是立刻调用该匿名函数，更加灵活常见

	go func() {
		fmt.Println("Hello goroutine222")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello goroutine333")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello goroutine444")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello goroutine555")
		wg.Done()
	}()

	wg.Wait() //阻塞作用：计数器为0时返回

}

func main() {
	//goroutine的调度是随机的
	testgroup()
}
