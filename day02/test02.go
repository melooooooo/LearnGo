package main

import (
	"fmt"
)

//func test01(n1 *int){
//	*n1=*n1+1
//}
//func main(){
//	var num1 int=1
//	test01(&num1)
//	fmt.Println(num1)
//}


func getsum(n1 int,n2 int) int{
	return n1+n2
}

//函数既然是一种数据类型，因此可以作为形式参数进行调用
func myfun(fun func(int,int)int,num1 int,num2 int) int{
	return fun(num1,num2)
}

func main(){
	a:=getsum
	//fmt.Printf("a的类型为%T,getsum的类型为%T\n",a,getsum)
	////a的类型为func(int,int) int,getsum的类型为func(int,int) int
	//res:=a(1,2)
	//fmt.Println("res=",res)

	res := myfun(a, 1, 2)
	fmt.Println("res=",res)


}

//支持对函数的返回值命名
func getSumSub(n1 int,n2 int)(sum int,sub int){
	sum=n1+n2
	sub=n1-n2
	return
}

//go支持可变参数 args是slice切片，可以通过args[index]可以访问到各个值
func sum(n1 int,args... int) int{
	res:=n1
	for i:=0;i<len(args);i++{
		res+=args[i]
	}
	return res
}

