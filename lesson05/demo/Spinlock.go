package main

//简单地实现一个自旋锁

type Spinlock struct {
	value bool
}

func (sl *Spinlock) lock() {
	//for语句体现自旋结构
	for {
		if !sl.value {
			sl.value = true
			return
		}
	}
}

func (sl *Spinlock) unlock() {
	sl.value = false
}
