package main

import {
	"fmt"

}

func hello(i,int){
	println("hello go:",fmt.println(i))
}
func main()  {
	for i := 0; i < 5; i++ {
		go func(j int){
			hello(j)
		}(i)
	}
	time.sleep(time.Second)
}