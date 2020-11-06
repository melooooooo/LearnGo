package main

import "fmt"

func main() {
	//var a = make([]string, 5, 10)
	//for i := 0; i < 10; i++ {
	//	a = append(a, fmt.Sprintf("%v", i))
	//}
	//fmt.Println(a)
	//fmt.Println(len(a))
	//str:="hello@atguigu"
	//arr1:=[]byte(str)
	//arr1[1]='a'
	//str=string(arr1)
	//println(str)
	//
	//arr2:=[]rune(str)
	//arr2[0]='北'
	//str=string(arr2)
	//println(str)

	//make函数的作用在使用slice,map,channel的过程中都要使用make进行初始化，然后才可以对他们进行操作
	//var b map[string]int
	//b=make(map[string]int,10)
	//b["沙河娜扎"]=100
	//fmt.Println(b)

	//var arr3=[2][3]int{{1,2,3},{4,5,6}}
	//for i,v :=range arr3{
	//	for j,v1:=range v{
	//		fmt.Printf("arr3[%v][%v]=%v\t",i,j,v1)
	//	}
	//}



	//map的增删改查
	cities:=make(map[string]string,10)
	cities["no1"]="beijing"
	cities["no2"]="tianjing"
	cities["no3"]="shanghai"
	fmt.Println(cities)
	fmt.Println(len(cities))

	delete(cities,"no1")
	//delete(cities,"no4")
	fmt.Println(cities)


	//map的增删改查
	val,ok:=cities["no1"]
	if ok{
		fmt.Printf("有no1 key 值为%v\n",val)
		fmt.Println(ok)
		fmt.Printf("val type=%T,val=%v",val,val)
	}else{
		fmt.Printf("没有no1 key\n")
		fmt.Println(ok,val)
		fmt.Println(val=="")
	}
}
