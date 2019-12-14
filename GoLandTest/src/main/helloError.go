package main

import "fmt"
import "errors"

//自定义的出错结构
type myError struct {
	arg int
	errMsg string
}

//实现Error接口
func (e *myError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.errMsg)
}

//两种出错
func error_test(arg int) (int, error)  {
	if arg < 0 {
		return -1, errors.New("Bad Arguments - negtive!")
	}else if arg > 256 {
		return -1, &myError{arg, "Bad Arguments - too large!"}
	}
	return arg*arg, nil
}

//defer语句  在 Go 函数中添加多个defer语句，当函数执行到最后时，
//这些 defer 语句会按照逆序执行（即最后一个defer语句将最先执行），最后该函数返回。
func deferTest() (result int) {
	defer func() {
		result = 12
	}()
	return 10
}

func foo() {
	panic(errors.New("I am a bug"))
	return
}

func deferTest2() (result int) {
	defer func() {
		if r := recover(); r!= nil {
			err := r.(error)
			fmt.Println("Cache Exception:", err)
		}
	}()
	foo()
	return 10
}

//相关的测试
func main(){
	for _, i := range []int{-1,4,1000} {
		if r, e := error_test(i); e != nil {
			fmt.Println("failed:", e)
		} else {
			fmt.Println("success:", r)
		}
	}
	fmt.Println("deferTest:")
	fmt.Println(deferTest())
	fmt.Println("deferTest2:")
	fmt.Println(deferTest2())
}