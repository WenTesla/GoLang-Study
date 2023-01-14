package main

import "fmt"

/*
编写一个程序来转换罗马数字（例如将 MCLX 转换成 1,160）。
使用映射加载要用于将字符串字符转换为数字的基本罗马数字。
例如，M 将是映射中的键，其值将为 1000。 使用以下字符串字符映射表列表：

M => 1000
D => 500
C => 100
L => 50
X => 10
V => 5
I => 1
如果用户输入的字母与上述列表中的不同，则打印一个错误。
*/
//初始化
var romanMap = map[rune]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}

func romanToArabic(s string) int {
	arabicVals := make([]int, len(s)+1)
	for index, digit := range s {
		if val, present := romanMap[digit]; present {
			arabicVals[index] = val
		} else {
			fmt.Printf("Error: The roman s %s has a bad digit: %c\n", s, digit)
			return 0
		}
	}

	total := 0

	for index := 0; index < len(s); index++ {
		if arabicVals[index] < arabicVals[index+1] {
			arabicVals[index] = -arabicVals[index]
		}
		total += arabicVals[index]
	}

	return total
}
func main() {
	//var x string
	//fmt.Scanf("%s", &x)
	//fmt.Printf(string(romanToArabic(x)))
	fmt.Println("MCLX is", romanToArabic("MCLX"))
	fmt.Println("MCMXCIX is ", romanToArabic("MCMXCIX"))
	fmt.Println("MCMZ is", romanToArabic("MCMZ"))
}
