package main

import "fmt"

func main(){
	str2:="hello北京"
	r:=[]rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Println("字符=%c\n",r[i])
		//fmt.Println("字符=%c\n",r[i])
	}
	
}
