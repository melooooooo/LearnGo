package main

import "fmt"

type Cat struct{
	Name string
	age int
}
func main(){
	var allChan =make(chan interface{},10)
	cat1:=Cat{"tom",18}
	cat2:=Cat{"tom~",180}
	allChan<-cat1
	allChan<-cat2
	allChan<-10
	allChan<-"jack"
	temp1:=<-allChan
	fmt.Printf("temp1=%v,temp1 type=%T\n",temp1,temp1)
	//在编译层面下面这个会报错，运行时才可以确定为main.Cat类型
	//fmt.Println(temp1.Name) 可以使用类型断言
	a:=temp1.(Cat)
	fmt.Println(a.Name)
	temp2:=<-allChan
	fmt.Printf("temp2=%v,temp2 type=%T\n",temp1,temp1)
	fmt.Println(temp1,temp2)

}
