package main

import "fmt"

//结构体中 %v 只输出所有的值
//%+v 先输出字段类型，再输出该字段的值
//%#v的区别 先输出结构的名字值，在输出结构体

//type A struct{
//	Name string
//	age int
//}

type B struct{
	Name string
	Score string
}

type C struct{
	//A
	B
}

func main(){
	var c C
	c.B.Name="tom"
	fmt.Println(c)
}
