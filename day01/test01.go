package main

import "fmt"

func main(){
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool=true
	var myChar byte='h'
	var str string
	var str string
	str=fmt.Sprintf("%d",num1)
	fmt.Printf("str type =%T,str=%q\n",str,str)
	str=fmt.Sprintf("%f",num2)
	fmt.Printf("str type =%T,str=%q\n",str,str)
	str=fmt.Sprintf("%t",b)
	fmt.Printf("str type =%T,str=%q\n",str,str)
	str=fmt.Sprintf("%c",myChar)
	fmt.Printf("str type =%T,str=%q\n",str,str)

}