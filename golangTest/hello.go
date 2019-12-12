package main

import (
    "fmt"
)

func main(){
	const (
		limit = iota
		top  = iota
	)
	const Pi  = iota  
	const x,y int = 1,3

	var a int
	var b string
	var c float64
	var d [5] int
	var e [] int
	var f * int
	var v1 int = 4
	var  v2 = 3
	v3 := 5
	v1,v2 = v2,v1

    fmt.Println("Hello,World!Are you OK?")
    fmt.Println(limit,top,Pi,x,y)
    fmt.Println(a,b,c,d,e,f,v1,v2,v3)
}
