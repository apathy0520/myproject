package main

import (
	"fmt"
)

func main(){
	//类型转换 resultOfType := Type(expression)
	//可以通过type关键字声明类型，如type StringsSlice []string 将[]string（string 类型的切片）声明为StringSlice类型
	x := int16(2345)
	y := int32(x)
	a := uint16(65000)
	b := int16(a)
	fmt.Println(x,y,a,b)
	//%T 格式化指令的作用是输出变量的类型
	fmt.Printf("type and value of a is: %T and %d\n", a, a)
	fmt.Printf("type and value of b is: %T and %d\n", b, b)

	var i interface{} = 99	//创建一个interface{}类型。其值为99
	var s interface{} = []string{"left", "right"}
	j := i.(int) //假设i是兼容int类型，并使用类型断言将其转换为int类型
	fmt.Printf("type and value of j is: %T and %d\n", j, j)

	if s, ok := s.([]string); ok {	//创建影子变量，if的作用域中覆盖了外部的变量s
		fmt.Printf("%T -> %q\n", s, s)
	}
}