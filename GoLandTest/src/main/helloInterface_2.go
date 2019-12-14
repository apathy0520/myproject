package main

import (
	"fmt"
	"strconv"
)

//在接口里也可以再嵌入接口。
//如果一个interface1作为interface2的一个嵌入字段，
//那么interface2隐式的包含了interface1里的方法。
//如下例子中, Interface2包含了Interface1的所有方法。

type Interface1 interface {
	Send()
	Receive()
}

type Interface2 interface {
	Interface1
	Close()
}

type Element interface {}
type List []Element

type Person struct {
	name string
	age int
}

//实现了fmt.Stringer接口
func (p Person) String() string {
	return "(name: "+ p.name + " - age: "+strconv.Itoa(p.age) + "years)"
}

func main() {
	list := make(List,4)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"Dennis", 70}
	list[3] = 1.223

	for index, element := range list {
		switch value := element.(type) {	//switch类型判断开关
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Printf("list[%d] is of a different type", index)
		}
	}
}