package main

import (
	"fmt"
	"sync"
)

//基于channel和sync实现go信号量

//信号量，是一种控制并发访问共享资源的同步机制，核心在于维护一个计数器，表示可用资源的数量
//当协程获取资源时，计数器-1，释放时+1，从而限制同时访问资源的协程数量

//使用带缓冲channel可以简单实现信号量，buffer容量即为信号量初始值，发送操作代表获取信号量，接收代表释放

type Semaphore chan struct{} //空结构体类型，无内存，作为信号相当合适

func (s Semaphore) Acquire() {
	s <- struct{}{}
}

func (s Semaphore) Release() {
	<-s
}

func main() {
	sem := make(Semaphore, 2)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem.Acquire()
			fmt.Printf("goroutine %d start conducting.. \n", id)
			sem.Release()
		}(i)
	}
	wg.Wait()

	//慢慢的，因为限制了访问数量
}
