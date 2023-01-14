package main

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	counts:=make(map[string]int)
	input:=bufio.NewScanner(os.Stdin)
	fmt.Println("请输入：")
	for input.Scan() {
		if input.Text() == "q"{
			break
		}
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}