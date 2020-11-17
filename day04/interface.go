package main

import (
	"fmt"
	"reflect"
)

func TypeJudge(item... interface{}){
	for index,value:=range item{
		switch value.(type) {
		case bool:
			fmt.Printf("第%d个参数的类型时bool类型,值为%v\n",index,value)
		case float32:
			fmt.Printf("第%d个参数的类型时float32类型,值为%v\n",index,value)
		case float64:
			fmt.Printf("第%d个参数的类型时float64类型,值为%v\n",index,value)
		case int,int32,int64:
			fmt.Printf("第%d个参数的类型时int类型,值为%v\n",index,value)
		case string:
			fmt.Printf("第%d个参数的类型时string类型,值为%v\n",index,value)
		default:
			fmt.Printf("第%d个参数的类型不确定,值为%v\n",index,value)
		}
		
	}
}
func main(){
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 =39
	var name string ="tom"
	var n5 int64=980
	address:="北京"
	n4:=300
	TypeJudge(n1,n2,n3,name,address,n4)
	fmt.Println(reflect.TypeOf(n4))
	fmt.Println(reflect.TypeOf(n3))
	fmt.Println(reflect.TypeOf(n5))
	fmt.Println(reflect.TypeOf(n5)==reflect.TypeOf(n4))
	fmt.Println(reflect.TypeOf(n3)==reflect.TypeOf(n4))

}
