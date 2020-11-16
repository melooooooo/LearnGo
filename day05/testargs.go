package main

import (
	"errors"
	"fmt"
)

func test(args... interface{}){
	fmt.Println(len(args))
}


func divide (a int, b int) (num int, err error){ //定义两个返回值
	//return a / b
	if b == 0 {
		err := errors.New("被除数不能为零！")
		return 0,err
	}
	return a / b, nil   //支持多个返回值
}

func main(){
	test(1,2,3,4)

}