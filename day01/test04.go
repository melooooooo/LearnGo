package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	//var str string ="hello,world!北京"
	//str1:=[]rune(str)
	//for i:=0;i<len(str1);i++{
	//	fmt.Printf("%c\n",str1[i])
	//}
	//
	//str2:="abc~ok上海"
	//for index,val:=range str2{
	//	fmt.Printf("index=%d,value=%c",index,val)
	//}
	//
	//fmt.Println(len(str))
	//for i:=0;i<len(str);i++{
	//	fmt.Println(str[i])
	//}
	//for index,value:=range(str){
	//	//fmt.Println("index=",index,"value=",value)
	//	fmt.Printf("index=%d,value=%c",index,value)
	//}
	//
	////使用for实现while...do循环
	//var i int =1
	//for{
	//	if i>10{
	//		break
	//	}
	//	fmt.Println("hello,world!",i)
	//	i++
	//}
	////利用for实现do...while循环
	//var j int =1
	//for{
	//	fmt.Println("hello,ok",j)
	//	j++
	//	if j>10{
	//		break
	//	}
	//}

	//rand.Seed(time.Now().Unix())
	var count int =0
	rand.Seed(time.Now().Unix())
	for{
		n:=rand.Intn(100)+1
		fmt.Println("n=",n)
		count++
		if n==99{
			fmt.Println("随机到99一共循环了",count,"次")
			break
		}
	}
}
