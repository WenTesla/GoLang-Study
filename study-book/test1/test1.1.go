package main

import (
    "fmt"
	"os"
	"strings"
)

func main() {
	// 法一
	fmt.Println(strings.Join(os.Args, " "));
	// 法二
	fmt.Println(os.Args[0])
	
}