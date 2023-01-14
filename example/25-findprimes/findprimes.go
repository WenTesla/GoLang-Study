package main

import "fmt"

/**
编写一个程序来查找小于 20 的所有质数。 
质数是大于 1 的任意数字，只能被它自己和 1 整除。
 “整除”表示经过除法运算后没有余数。 与大多数编程语言一样，Go 还提供了一种方法来检查除法运算是否产生余数。 
 我们可以使用模数 %（百分号）运算符。
*/
func findprimes(number int) bool {
	for i := 2; i < number; i ++ {
        if number % i == 0 {
			return false
        }
    }

	if number > 1{
		return true
	} else {
	    return false
	}
}

func main() {
    fmt.Println("Prime numbers less than 20:")

    for number := 2; number <= 20; number++ {
        if findprimes(number) {
            fmt.Printf("%v ", number)
        }
    }
}