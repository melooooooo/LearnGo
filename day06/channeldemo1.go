package main

import "fmt"

func main(){
	var c chan int=make(chan int,3)
	fmt.Printf("c的值为%v c本身的地址为%v",c,c)
	c<-10
	num:=11
	c<-num
	c<-50

	fmt.Printf("channel c len=%v,cap=%v",len(c),cap(c))

	num1:=<-c
	num2:=<-c
	num3:=<-c
	//如果试图向一个空的channel中取数据的话会报deadlock的错误
	//num4:=<-c
	//fmt.Println(num4)
	fmt.Println("num1=",num1,"num2=",num2,"num3=",num3)
}
