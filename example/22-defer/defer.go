package main

import (
	"fmt"
)

func main()  {
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	defer f(10)
	defer fmt.Println("5")
	defer fmt.Println("6")
	defer fmt.Println("7")
	defer fmt.Println("8")

	
}
func f(s int){
	fmt.Println(s)
}