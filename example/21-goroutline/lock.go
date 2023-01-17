package main

import (
	"sync"
	"time"
)

var (
	x    int64
	lock sync.Mutex
)

func calcWithLock() {
	for i := 0; i < 2000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}

func calcWithoutLock() {
	for i := 0; i < 2000; i++ {
		x += 1
	}
}

func main() {
	x = 0
	for i := 0; i < 5; i++ {
		go calcWithLock()
	}
	//隔绝上面的协程，上面的全部进行完成后再输出
	time.Sleep(time.Second)
	println("开启锁", x)
	x = 0
	for i := 0; i < 5; i++ {
		go calcWithoutLock()
	}
	time.Sleep(time.Second)
	println("不开启锁", x)
}
