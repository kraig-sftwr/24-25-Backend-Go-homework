package main

import "fmt"

//先讲一下go channel

func main() {
	cha := make(chan int, 100)    //声明一个channel类型（可以发送和接收）
	cha2 := make(chan<- int, 100) //只能发送给cha2
	//var cha3 <-chan int //只能从cha3接收

	//cha4 := make(chan int, 100) //make函数初始化cha4，容量为100，代表cha1的缓冲区
	//如果不设置容量（或容量设置为0），说明channel无缓存，只有发送者和接收者都准备好后才会开始通讯
	//设置缓存后可能不发生阻塞，只有buffer满后发送阻塞，buffer空后接收阻塞
	//nil channel不发生通讯
	//channel作为FIFO队列

	//var cha4 chan int //nil channel

	cha <- 10 //send语句
	cha <- 11

	v1, ok := <-cha //channel接收支持multi
	v2 := <-cha

	cha2 <- 12

	fmt.Println("cha: ", v1, v2, ok)
	fmt.Println("cha2: ", cha2) //直接输出cha2会发生什么呢...---地址！
	//我草这个go channel太好用了你们知道吗

}
