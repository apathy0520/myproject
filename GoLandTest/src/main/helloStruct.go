package main

import "fmt"

//组合
type Count int	//创建自定义类型Count

func (count *Count) Increment()	{ *count++ }	//Count类型的方法
func (count *Count) Decrement()	{ *count-- }
func (count Count) IsZero()	bool { return count == 0 }

type Part struct{	//基于结构体创建自定义类型Part
	stat string
	Count //匿名字段
}

//先定义了一个 Count 类，然后定义了一个 Part 类。
//Part 类“继承”了 Count 类的所有成员属性和方法并重写了 IsZero()方法。
//同时在重写 IsZero()方法时调用了 Count 类的 IsZero()方法。
//需要注意的是，若此时 Part 的对象通过part.IsZero()调用 IsZero() 方法时，
//实际调用的是 Part 重写过后的 IsZero()方法，而不是基类 Count 的 IsZero()方法，
//若想调用 Count 类的 IsZero()方法则要使用part.Count.IsZero(),而调用没有重写的 Decrement()方法时，
//使用part.Decrement()和part.Count.Decrement()效果是一样的。

func (part Part) IsZero() bool {	//覆盖了匿名字段Count的IsZero（）方法
	//part.IsZero()		//调用本身的IsZero()方法，而不是Count的
	return part.Count.IsZero() && part.stat == ""	//调用了匿名字段Count的方法
}

func (part Part) String() string {	//定义String（）方法，自定义了格式化指令%v的输出
	return fmt.Sprintf("<<%s, %d>>", part.stat, part.Count)
}

//定义的类型Count中，*Count方法集是Increment(), Decrement()和IsZero()，
//Count的值的方法集是IsZero()。
//但是因为Count类型的是可寻址的，所以我们可以使用Count的值调用全部的方法。
//另外如果结构体的字段也有方法，我们也可以直接通过结构体访问字段中的方法

func main(){
	var i Count = -1
	fmt.Printf("Start \"Count\" test:\nOrigin value of count: %d\n", i)
	i.Increment()
	fmt.Printf("Value of count after increment: %d\n", i)
	fmt.Printf("Count is zero t/f? : %t\n\n", i.IsZero())
	fmt.Println("Start: \"Part\" test:")
	part := Part{"232", 0}
	fmt.Printf("Part: %v\n", part)
	fmt.Printf("Part is zero t/f? : %t\n",part.IsZero())
	fmt.Printf("Count in Part is zero t/f?: %t\n", part.Count.IsZero())	//尽管覆盖了匿名字段的方法，但还是可以访问
}