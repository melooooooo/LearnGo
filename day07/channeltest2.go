package main

import "fmt"

func main(){
	var intChan chan<- int
	intChan=make(chan int,3)
	intChan<- 20

	fmt.Println("intChan=",intChan)

	var chan3 <-chan int
	num2:=<-chan3
	fmt.Println(num2)

	fmt.Println()

}
