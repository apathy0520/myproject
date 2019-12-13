package main

import(
	"fmt"
)

func main(){

	//s += t	将字符串 t 追加到 s 末尾
	//s + t	将字符串 s 和 t 级联
	//s[n]	从字符串 s 中索引位置为 n 处的原始字节
	//s[n:m]	从位置 n 到位置 m-1 处取得的字符（字节）串
	//s[n:]	从位置 n 到位置 len(s)-1 处取得的字符（字节）串
	//s[:m]	从位置 0 到位置 m-1 处取得的字符（字节）串
	//len(s)	字符串 s 中的字节数
	//len([]rune(s))	字符串 s 中字符的个数，可以使用更快的方法 utf8.RuneCountInString()
	//[ ]rune(s)	将字符串 s 转换为一个 unicode 值组成的串
	//string(chars)	chars 类型是[]rune 或者[]int32, 将之转换为字符串
	//[ ]byte(s)	无副本的将字符串 s 转换为一个原始的字节的切片数组，不保证转换的字节是合法的 UTF-8 编码字节
	t0 := "\u6B22\u8FCE\u6765\u5230"
	t1 := "\u5B9E\u9A8C\u697C"
	t2 := t0 + t1
	for index, char := range t2 {
	    fmt.Printf("%-2d   %U    '%c'   %X    %d\n",
	        index, char, char, []byte(string(char)), len([]byte(string(char))))
	}
	fmt.Printf("length of t0: %d, t1: %d, t2: %d\n",len(t0), len(t1), len(t2))
	fmt.Printf("content of t2[0:2] is: %X\n", t2[0:2])
}
