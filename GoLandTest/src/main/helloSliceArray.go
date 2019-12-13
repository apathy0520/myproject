package main

import (
	"fmt"
)

func main() {
	test()
}
func test(){
	a := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("len and cap of array %v is: %d and %d\n", a, len(a), cap(a))
	fmt.Printf("item in array: %v is:", a)
	for _, value := range a {
		fmt.Printf("% d", value)
	}

	fmt.Println()

	s1 := a[3:6]
	fmt.Printf("len and cap of slice: %v is: %d and %d\n", s1, len(s1), cap(s1))
	fmt.Printf("item in slice: %v is:", s1)
	for _, value := range s1 {
		fmt.Printf("% d", value)
	}

	fmt.Println()

	s1[0] = 456
	fmt.Printf("item in array changed after changing slice: %v is:", s1)
	for _, value := range a {
		fmt.Printf("% d", value)
	}

	fmt.Println()

	//内置函数make()用于创建切片、映射和通道。当用于创建一个切片时，
	//它会创建一个隐藏的初始化为零值的数组，然后返回一个引用该隐藏数组的切片。
	//该隐藏的数组与 Go 语言中的所有数组一样，都是固定长度，如果使用第一种语法创建，那么其长度为切片的容量capacity；
	//如果是第二种语法，那么其长度记为切片的长度length。一个切片的容量即为隐藏数组的长度，而其长度则为不超过该容量的任意值。
	//另外可以通过内置的函数append()来增加切片的容量。

	s2 := make([]int, 10, 20)
	s2[4] = 5
	fmt.Printf("len and cap of slice: %v is: %d and %d\n", s2, len(s2), cap(s2))
	fmt.Printf("item in slice %v is:", s2)
	for _, value := range s2 {
		fmt.Printf("% d", value)
	}

	fmt.Println()
}