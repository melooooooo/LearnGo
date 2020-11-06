package main

import "fmt"

func main(){
	//strslice:=[]string{}
	//println(cap(strslice))
	//for i := 0; i < 10; i++ {
	//	strslice=append(strslice,"test")
	//	println(cap(strslice))
	//}

	//a:=make([]int,2,10)
	//fmt.Println(a)
	//fmt.Println(len(a))
	//fmt.Println(cap(a))
	//切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针，切片的长度和切片的容量
	//for index,value:=range a{
	//	fmt.Println(index,value)
	//}
	//从切片中删除元素
	//sli:=[]int{1,2,3,4,5}
	//slice中添加另一个切片中的元素,需要加上...
	//表示对slice打散进行传递
	//sli=append(sli[:2],sli[3:]...)
	//fmt.Print(sli)
	var array =[5]int{1,2,3,4,5}
	var slice1=array[:]
	slice1=append(slice1[0:1],slice1[2:]...)
	fmt.Println(slice1)
	fmt.Println(array)

	//1.go语言中的copy函数
	//不同类型的切片无法复制
	//2.如果s1的长度大于s2的长度，将s2中对应位置上的值替换s1中对应位置的值
	//3.如果s1的长度小于s2的长度，多余的将不做替换



}
