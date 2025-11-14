package main

import (
	"fmt"
	"sync"
	"time"
)

//简单地实现一个互斥锁

//sync.Mutex中有两个方法Lock和Unlock

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Dec() {
	c.mu.Lock()
	c.count--
	c.mu.Unlock()
}

func main() {
	counter := Counter{}
	var wg sync.WaitGroup

	//启动多个goroutine同时计数，是否可以有效提高执行效率？

	st := time.Now()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()
	fmt.Println("Final count: ", counter.count)

	fmt.Println(time.Since(st))
	//--好慢！其实使用mutex的目标并非为了提升执行效率，而是为了保证复杂数据的正确性，有效避免数据竞争
}
