package other

import "errors"

func Add(a int,b int) (num int){
	return a + b
}

//函数多返回值
func Divide(a float64, b float64) (num float64, err error){
	if b == 0 {
		err = errors.New("被除数不能为0！")
		return -1, err
	}
	err = errors.New("正确！")
	return a / b, err
}