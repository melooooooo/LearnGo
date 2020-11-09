package main

import (
	"fmt"
)

type person struct{
	Name string
	Age int
	Scores [5]float64
	ptr *int
	slice []int
	map1 map[string]string
}

func main(){
	var p1 person
	fmt.Printf("type of p1=%T\n",p1)
	fmt.Println(p1)
	fmt.Println(p1.Name=="")
	fmt.Println(p1.Age==0)
	//fmt.Println(p1.Scores==[0,0,0,0,0])
	//fmt.Println(p1.Scores==[5]float64(0,0,0,0,0))
	fmt.Println(p1.slice==nil)
	fmt.Println(p1.map1==nil)
	fmt.Println(p1.ptr==nil)

	//使用slice的过程中首先要make
	p1.slice=make([]int,10)
	p1.slice[0]=100
	//使用map的过程中,首先要make
	p1.map1=make(map[string]string)
	p1.map1["key1"]="tom"

	fmt.Println(p1)


	//不同结构体的变量相互独立，互不影响，一个结构体变量字段的改变，不影响另外一个，结构体是值类型
	//结构体的声明方式
	//1.直接声明
	//var p2 Person
	//
	//var person Person=Person{}
	//
	//var person *Person =new(Person)
	//
	//var person *Person=&Person{}

	//第三种和第四种方式返回的是结构体的指针
	//结构体指针访问字段的标准方式是(*结构体指针).字段名 但是go语言底层做了简化也支持结构体指针.字段名
	//go编译器对person.Name做了转化(*person).Name



}
