package main

import "fmt"

func cal(n1 float64,n2 float64,operator byte) float64{
	var res float64
	switch operator {
	case '+':
		res=n1+n2
	case '-':
		res=n1-n2
	case '*':
		res=n1*n2
	case '/':
		res=n1/n2
	default:
		fmt.Println("操作符号错误！")
	}
	return res
}
func main(){
	//var res float64
	var n1 float64
	var n2 float64
	var op byte
	fmt.Println("请输入数字1，数字2，和操作符:")
	fmt.Scanf("%f %f %c",&n1,&n2,&op)
	result:=cal(n1,n2,op)
	fmt.Println()
	fmt.Println("最后的结果是",result)
}
