package main

import (
	"fmt"
	"other"	//导入其他包
)

func main(){
	c := other.Add(5,2)  //函数调用
	fmt.Println(func(j, k, l float64) float64 { return j * k / l }(5.2,6.4,3.3))	//匿名函数1
	t := func (q,w,e int) bool { return q*w<e }(1,2,3)	//匿名函数2
	a, b := other.Divide(55,2)	//函数多返回值
	x, y := other.Divide(4,0)
	fmt.Println(c,a,b,x,y,t)
}