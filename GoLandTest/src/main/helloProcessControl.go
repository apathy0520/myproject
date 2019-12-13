package main

import (
	"fmt"
)

func main(){
	myfunc1()
	myfunc2()
	myfunc3(5,-17.98,"AIDEN",nil,true,complex(1,1))
}

//跳转语句
func myfunc1(){
	i := 0;
	THIS:  //定义一个THIS标签
	fmt.Println(i)
	i++
	if i < 5 {
		goto THIS  		//跳转到标签处
	}
}
//for语句
func myfunc2(){
	i := 0
	flag := false
	for flag == false {  // while循环，在Go语言中没有while关键字
		i++
		if i == 5 {
			flag = true
		}
	}
	fmt.Println(i,flag)
	str := "asdfghjkl"
	for index, value := range str {		// 迭代字符串
		fmt.Println(index,value)
	}
	//for {
	//	fmt.Println(i)			//无限循环
	//}
}
//switch语句
func myfunc3(items ...interface{}){
	for i, x := range items {
		switch x := x.(type) {	//创建了影子变量
		case bool:
			fmt.Printf("param #%d is a bool, value: %t\n", i, x)
		case float64:
			fmt.Printf("param #%d is a float64, value: %f\n", i, x)
		case int, int8, int16, int32, int64:
			fmt.Printf("param #%d is a int, value: %d\n", i, x)
		case uint, uint8, uint16, uint32, uint64:
			fmt.Printf("param #%d is a uint, value: %d\n", i, x)
		case nil:
			fmt.Printf("param #%d is a nil\n", i)
		case string:
			fmt.Printf("param #%d is a string, value: %s\n", i, x)
		default:
			fmt.Printf("param #%d's type is unknow\n", i)
		}
	}
}