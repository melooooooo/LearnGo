package main

import (
	"fmt"
	"math"
)

func main(){
	//fmt.Println("请输入你的姓名")
	//var name string
	//fmt.Scanln(&name)
	//fmt.Println("请输入你的成绩")
	//var score int
	//fmt.Scanln(&score)
	//fmt.Printf("你的姓名是%v,你的成绩是%v",name,score)
	var a float64=2.0
	var b float64=4.0
	var c float64=2.0
	//这里很可能会出现浮点数的精度的问题
	m:=b*b-4*a*c
	if m>0 {
		x1:=(-1*b+math.Sqrt(m))/2*a
		x2:=(-1*b+math.Sqrt(m))/2*a
		fmt.Printf("x1=%f,x2=%f",x1,x2)
	}else if m==0 {
		x1:=(-1*b+math.Sqrt(m))/2*a
		fmt.Printf("x1=%f",x1)
	}else{
		fmt.Println("最后无解")
	}
}
