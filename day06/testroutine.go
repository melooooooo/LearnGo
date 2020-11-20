package main

import "fmt"

func main(){
	test:=make(chan int,1)
	//close(test)
	//test<-1
	v,ok:=<-test
	fmt.Println(v,ok)
}
