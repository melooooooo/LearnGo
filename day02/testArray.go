package main

import "fmt"

func main(){
	//var score [5]int
	//for i := 0; i < len(score); i++ {
	//	fmt.Printf("请输入第%v个元素的值\n	",i)
	//	fmt.Scanln(&score[i])
	//}
	//
	//for i := 0; i < len(score); i++ {
	//	fmt.Printf("score[%v]=%v",i,score[i])
	//}
	//var numArr01=[3]int{5,6,7}
	//var numArr01 [3]int={1,2,3}
	//var numArr02 [3]int={5,6,7}
	//var numArr03 [3]int={7,8,9}
	//var numArr04=[...]int{1,2,3}
	//var numArr02=[3]int{1,2,3}
	//for _,value :=range numArr02{
	//	println(value)
	//}
	//fmt.Println(numArr02)
	//var arr=[5]int{1,2,3,4,5}
	//var slice=arr[:]
	//fmt.Println("arr=",arr)
	//fmt.Println("slice=",slice)
	//fmt.Println("slice length=",len(slice))
	//fmt.Println("slice capacility=",cap(slice))
	var strslice=[]string{"tom","jack","mary"}
	fmt.Println("strslice=",strslice)
	fmt.Println("strslice size=",len(strslice))
	fmt.Println("strslice cap=",cap(strslice))
	strslice= append(strslice, "rookie")
	println(cap(strslice))
	strslice=append(strslice, "tomcat")
	println(cap(strslice))
}
