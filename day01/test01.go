package main

import (
	"fmt"
	"strconv"
)

func main(){
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool=true
	var myChar byte='h'
	var str string
	//1.第一种方式
	str=fmt.Sprintf("%d",num1)
	fmt.Printf("str type =%T,str=%q\n",str,str)
	str=fmt.Sprintf("%f",num2)
	fmt.Printf("str type =%T,str=%q\n",str,str)
	str=fmt.Sprintf("%t",b)
	fmt.Printf("str type =%T,str=%q\n",str,str)
	str=fmt.Sprintf("%c",myChar)
	fmt.Printf("str type =%T,str=%q\n",str,str)
	//2.第二种方式
	//%q 返回的是该值对应的单引号括起来的go语法字符字面值
	res1 := strconv.FormatInt(int64(num1), 10)
	fmt.Printf("res1 type=%T,res1=%q",res1,res1)
	res2 := strconv.FormatFloat(num2, 'f', 10, 64)
	fmt.Printf("res2 type=%T,res2=%q",res2,res2)

	var num5 int64 = 4567
	result:=strconv.Itoa(int(num5))
	fmt.Printf("str type=%T,str=%q",result,result)

	var boolstr string="true"
	res, _ := strconv.ParseBool(boolstr)
	//%v	值的默认格式表示
	fmt.Printf("res type=%T,type=%v",res,res)

	var numstr="123456"
	resnum, _ := strconv.ParseInt(numstr, 10, 64)
	fmt.Printf("res type=%T,value=%v",resnum,resnum)

	var floatstring="123.45"
	resfloat, _ := strconv.ParseFloat(floatstring, 64)
	fmt.Printf("res type=%T,res value=%v",resfloat,resfloat)

	//string转成基本数据类型的注意事项：在将String类型转成基本数据类型的时候，要确保String类型的数据能够转成有效的数据
	//如果不能转成有效的数据类型的话，golang将会把原本设定是整型的转为0，float->0,bool->false

	var i int = 10
	fmt.Println("i的地址是",&i)

	var ptr *int=&i
	fmt.Printf("ptr=%v",ptr)
	fmt.Printf("ptr=%v",&ptr)
	fmt.Printf("ptr指向的值=%v",*ptr)
	//golang中没有public,private等关键字，首字母大写是公开的，首字母小写是私有的
	//在goLang中自增和自减操作只能单独使用，而不能结合表达式进行使用
	//在golang中的自增自减操作符号只能放在后面，不能放在前面

}