package main
//
//import "fmt"
//
//func sum(n1 int,n2 int) int{
//	//在进行defer语句的代码放到栈中时，也会将相关的值拷贝同时放入栈中
//	defer fmt.Println("ok1 n1=",n1)
//	defer fmt.Println("ok2 n2=",n2)
//	n1++
//	n2++
//	res:=n1+n2
//	fmt.Println("ok3 res=",res)
//	return res
//}
//func main(){
//	res:=sum(10,20)
//	fmt.Println(res)
//}