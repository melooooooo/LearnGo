package main

import "fmt"

type Box struct{
	len float64
	width float64
	height float64
}
func (box *Box) getVolumn() float64{
	return box.len * box.height*box.width
}
func main(){
	//对于方法(如struct的方法),接受者为值类型时，可以直接用指针类型的变量调用,反过来同样也一样
	//对于普通的函数，接受者为值类型时，不能将指针类型的数据直接传递,反之亦然
	var box Box
	box.len=1.1
	box.width=2.0
	box.height=3.0
	volumn:=box.getVolumn()
	fmt.Println(volumn)



}
