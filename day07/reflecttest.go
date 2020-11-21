package main

import (
	"fmt"
	"reflect"
)

type Student struct{
	Name string
	Age int
}

func reflecttest01(b interface{}){
	//1.先获取到reflect.Type
	rType:=reflect.TypeOf(b)

	fmt.Println("rType=",rType)
	//2.获取到reflect.Value
	rValue:=reflect.ValueOf(b)
	n2:=2+rValue.Int()
	fmt.Println("n2=",n2)
	fmt.Printf("rVal=%v rValue Type=%T\n",rValue,rValue)//rVal=20 rValue Type=reflect.Value

	//3.下面我们将rValue转换为interface{}类型
	iV:=rValue.Interface()
	num2:=iV.(int)
	fmt.Println("num2=",num2)
	fmt.Println("num2的类型为",reflect.TypeOf(num2))
}

//针对结构体的反射
func reflecttest02(b interface{}){
	//1.先获取到reflect.Type
	rType:=reflect.TypeOf(b)
	fmt.Println("rType=",rType)
	//2.获取reflectValue
	rVal:=reflect.ValueOf(b)
	fmt.Println(rVal)
	iV:=rVal.Interface()
	stu,ok:=iV.(Student)
	if ok{
		fmt.Println("stu.Name=",stu.Name)
	}
}

func main(){
	var num int=20
	reflecttest01(num)
	stu:=Student{
		"tom",
		18,
	}
	reflecttest02(stu)
}
