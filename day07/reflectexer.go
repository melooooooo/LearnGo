package main

import (
	"fmt"
	"reflect"
)


func test(){
	var str string="tom"
	fs:=reflect.ValueOf(&str)
	fs.Elem().SetString("jack")
	fmt.Println("str=",str)
}

func main(){

	//var v float64=1.2
	//value := reflect.ValueOf(v)
	//t:=reflect.TypeOf(v)
	//fmt.Println("值为",value,"类型为",t)
	//i := value.Interface()
	//fmt.Printf("i=%v,type=%T",i,i)
	//f,ok := i.(float64)
	//if ok{
	//	fmt.Println("i的类型为float64")
	//	fmt.Println(f)
	//}
	test()
}
