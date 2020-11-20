package main

import "fmt"

func main(){
	intChan:=make(chan int,3)
	intChan<-100
	intChan<-200
	close(intChan)
	fmt.Println("okok~")
	//close了之后不能再写入数据
	n1:=<-intChan
	fmt.Println(n1)
	intChan2:=make(chan int,100)
	for i := 0; i < 100; i++ {
		intChan2<-i*2
	}
	close(intChan2)
	//如果管道没有关闭的话就对其进行range遍历就会出现deadlock错误
	for i:=range intChan2{
		fmt.Println("i=",i)
	}
}





