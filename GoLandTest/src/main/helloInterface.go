package main

import "fmt"

type Human struct {	//结构体
	name string
	age int
	phone string
}

//Human实现SayHi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n",h.name, h.phone)
}

//Human实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...",lyrics)
}

type Student struct {
	Human //匿名字段
	school string
	loan float32
}

type Employee struct {
	Human //匿名字段
	company string
	money float32
}

//Employee重载Human的SayHi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n",e.name, e.company, e.phone)
}

//Interface Men被Human，Student和Employee实现
//因为三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike",25,"222-222-xxx"},"MIT",0.00}
	paul := Student{Human{"Paul",26,"111-222-xxx"},"Harvard",100}
	sam := Employee{Human{"Sam",36,"444-222-xxx"},"Golang Inc",1000}
	tom := Employee{Human{"Tom",38,"222-444-xxx"},"Thing Ltd",5000}

	//定义Men类型的变量i
	var i Men

	//i能存储Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	//i也能存储Employee
	i = tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	//这三个都是不同类型的元素，但他们实现了interface同一个接口
	x[0],x[1],x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}

}