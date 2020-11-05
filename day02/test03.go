package main

import "fmt"

//如果一个文件中包含全局变量的定义，init函数和main函数，那么执行的顺序是全局变量定义->init函数->main函数
//init函数的主要作用就是做一些初始化的工作
func init(){
	fmt.Println("init...")
}

var age=test()

func test() int{
	fmt.Println("test()")
	return 90
}

func main(){
	fmt.Println("main()...age=",age)
}