package main

import "fmt"

type Goods struct{
	Name string
	Price float64
}

type Brand struct{
	Name string
	Address string
}
type TV struct{
	Goods
	Brand
}

type TV2 struct{
	*Goods
	*Brand
}

//使用嵌套匿名结构体
func main(){
	tv:=TV{Goods{"电视机001",5000},Brand{"haier","山东"}}
	tv2:=TV2{&Goods{"电视机002",4000},&Brand{"创维","河南"}}
	fmt.Println(tv)
	fmt.Println(tv2)
}
