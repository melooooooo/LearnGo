package main

import (
	"fmt"
)


func main(){
	x := func (a, b, c int) bool {
		return (a * b < c)
	}
	fmt.Println(x(1,2,3))
}


//func (a, b, c int) bool {
//	return (a * b < c)
//} (1, 2, 3)
