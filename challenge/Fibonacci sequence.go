/*你将编写一个程序来计算某个数字的斐波纳契数列。
你将编写一个函数，它返回一个包含按斐波纳契数列排列的所有数字的切片，而这些数字是通过根据用户输入的大于 2 的数字计算得到的。
让我们假设小于 2 的数字会导致错误，并返回一个 nil 切片。

请记住，斐波那契数列是一个数字列表，其中每个数字是前两个斐波那契数字之和。
例如，数字 6 的序列是 1,1,2,3,5,8，
数字 7 的序列是 1,1,2,3,5,8,13，
数字 8 的序列是 1,1,2,3,5,8,13,21，
以此类推。
*/

package main

import "fmt"

func fibonacci(n int) []int {
	if n < 2 {
		return make([]int, 0)
	}
	nums := make([]int, n)
	//初始化
	nums[0] = 1
	nums[1] = 1
	for i := 2; i < n; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}
	return nums
}

func main() {
	//定义输入的数据
	var num int
	fmt.Print("What's the Fibonacci sequence you want? ")
	fmt.Scanln(&num)
	fmt.Println("The Fibonacci sequence is:", fibonacci(num))
}
