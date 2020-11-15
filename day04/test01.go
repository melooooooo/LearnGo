package main

import "fmt"

func main(){
	//定义一个空的借口
	var x interface{}

	//var b2 float64="1.1"
	var c3="string"
	x=c3
	if res,ok := x.(float64);ok{
		fmt.Println("convert success!")
		fmt.Printf("y的类型时%T 值为%v\n",res,res)
	}else{
		fmt.Println("convert failed")
		fmt.Printf("res=%v,res type=%T",res,res);//如果转换失败，则会返回对应类型的零值
	}
	fmt.Println("执行接下来的代码...")
}

