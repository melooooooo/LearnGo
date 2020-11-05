package main

import "fmt"

func main(){
	//go语言中支持匿名函数，匿名函数
	res1:=func(n1 int,n2 int)int{
		return n1+n2
	}(10,20)

	fmt.Println(res1)

}
