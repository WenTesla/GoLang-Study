package main

import (
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func foo1() {
	log.Println("entry foo1")
	time.Sleep(5 * time.Second)
	wg.Done()
	log.Println("exit foo1")
}

func foo2() {
	log.Println("entry foo2")
	time.Sleep(2 * time.Second)
	wg.Done()
	log.Println("exit foo2")
}

func main() {
	log.Println("entry main")

	wg.Add(1)
	go foo1()

	wg.Add(1)
	go foo2()

	log.Println("wg.Wait()")
	wg.Wait()

	log.Println("exit main")
}
