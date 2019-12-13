package main

import(
	"fmt"
	"unsafe"
)

func main(){
	//数据类型：整型
	a := 12
	fmt.Println("length of a: " , unsafe.Sizeof(a))
	var b int = 12
	fmt.Println("length of b(int): " , unsafe.Sizeof(b))
	var c int8 = 12
	fmt.Println("length of c(int8): " , unsafe.Sizeof(c))
	var d int16 = 12
	fmt.Println("length of d(int16): " , unsafe.Sizeof(d))
	var e int32 = 12
	fmt.Println("length of e(int32): " , unsafe.Sizeof(e))
	var f int64 = 12
	fmt.Println("length of f(int64): " , unsafe.Sizeof(f))
	//数据类型：布尔型
	var aa bool
	aa = true
	bb := (2 == 3)
	fmt.Println(aa,bb)
	// var bbb bool
	// bbb = 1  //编译错误
	// fmt.Println(bbb)
	// bbb = bool(1)  //编译错误
	// fmt.Println(bbb)

	//数据类型：字符串    \\:表示反斜线  ':单引号 ":双引号 \n:换行符 \uhhh:4个16进制数字给定的Unicode字符
	t1 := "\"hello\""  
	t2 := `"hello"`
	t3 := "\u6B22\u8FCE"
	fmt.Println(t1,t2,t3)
}