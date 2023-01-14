// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type nameAndCount struct {
	filename []string
	count    int
}

func main() {
	//定义结构体 其中键值对为 出现的文本-结构体
	counts := make(map[string]*nameAndCount)
	files := os.Args[1:]
	if len(files) == 0 {
		print("files的长度为0")
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	// 输出
	for k, v := range counts {
		if v.count > 1 {
			fmt.Printf("\n 出现的次数：%d 序号：%s 文件名称：%v\n", v.count, k, v.filename)
		}
	}
}

func countLines(f *os.File, counts map[string]*nameAndCount) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		key := input.Text()
		print("当前扫描的值为：" + input.Text() + "\n")
		//如果存在值，追加
		if _, ok := counts[key]; ok {
			counts[key].count++
			counts[key].filename = append(counts[key].filename, f.Name())
		} else {
			//不存在值
			//手动添加值
			counts[key] = &nameAndCount{
				filename: make([]string, 1), //手动创建数值
				count:    1,
			}
			counts[key].filename[0] = f.Name()
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
