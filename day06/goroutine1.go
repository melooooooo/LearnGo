package main

import (
	"fmt"
)

func writeData(intChan chan int){
	//time.Sleep(time.Second*10)
	for i := 0; i < 10; i++ {
		intChan<-i
		fmt.Println("write data",i)
	}
	close(intChan)
}

func readData(intChan chan int,exitChan chan bool){
	for{
		v,ok:=<-intChan
		if !ok{
			break
		}
		fmt.Printf("read data 读到数据=%v\n",v)
	}
	exitChan<-true
	close(exitChan)
}

func main(){
	//test:=make(chan int,1)
	//close(test)
	//test<-1
	//v,ok:=<-test
	//fmt.Println(v,ok)
	intChan:=make(chan int,10)
	exitChan:=make(chan bool,1)
	go writeData(intChan)
	go readData(intChan,exitChan)
	for{
		_,ok:=<-exitChan
		if !ok{
			break
		}
	}
}
