package main

import (
	"fmt"
	"sync"
	"time"
)

var myMap map[int]int=make(map[int]int,10)
var lock sync.Mutex
func test(n int){
	res:=1
	for i := 1; i <= n; i++ {
		res*=i
	}
	lock.Lock()
	myMap[n]=res
	lock.Unlock()
}
func main(){
	for i := 0; i < 20; i++ {
		go test(i)
	}
	time.Sleep(time.Second)

	for i,v:=range myMap{
		fmt.Printf("map[%v]=%v",i,v)
	}
}


