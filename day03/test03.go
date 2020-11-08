package main

import "fmt"

type A struct{
	Num int
}

//表明(a A)体现test方法是和A类型绑定的

func (a A) test(){
	fmt.Println(a.Num)
}
func main(){
	var a A
	a.Num=100
	a.test()
}

