package main

import "fmt"

func PutNum(intChan chan int){
	for i := 0; i < 100; i++ {
		intChan <- i
	}
	close(intChan)
}

func PrimeNum(intChan chan int,primeChan chan int,exitChan chan bool){
	var flag bool
	for{
		num,ok:=<-intChan
		if !ok{
			break
		}
		flag=true
		for i := 2; i < num; i++ {
			if num%i==0{
				flag=false
				break
			}
		}
		if flag {
			primeChan<-num
		}
	}
	fmt.Println("有一个primeNum协程因为取不到数据，退出")
	exitChan <-true
}

func main(){
	intChan:=make(chan int,100)
	primeChan:=make(chan int,200)
	exitChan:=make(chan bool,4)
	go PutNum(intChan)
	for i := 0; i < 4; i++ {
		go PrimeNum(intChan,primeChan,exitChan)
	}
	go func(){
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	for{
		res,ok:=<-primeChan
		if !ok{
			break
		}
		fmt.Printf("素数=%d",res)
	}
	fmt.Println("线程退出")
}
