package main

import "fmt"

func main(){
	//scoremap:=make(map[string]int,10)
	//scoremap["zhangsan"]=90
	//scoremap["xiaoming"]=100
	//v,ok:=scoremap["zhangsan"]
	//if(ok){
	//	fmt.Println(v)
	//}else{
	//	fmt.Println("not foundß")
	//}
	//for key,value := range scoremap{
	//	fmt.Println(key,value)
	//}
	//for key := range scoremap{
	//	fmt.Println(key)
	//}
	var mapslice=make([]map[string]string,3)
	for index,value:=range mapslice{
		fmt.Printf("index:%v value=%v",index,value)
	}
		fmt.Println("after init")
		mapslice[0]=make(map[string]string,10)
		mapslice[0]["name"]="zhangsan"
		mapslice[0]["password"]="123456"
		mapslice[0]["address"]="沙河"
		for index,value:=range mapslice{
			fmt.Printf("index:%d value:%v\n",index,value)
		}
}
