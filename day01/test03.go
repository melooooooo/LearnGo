package main

import "fmt"

func main(){
	//var height int32
	//var money float32
	//var handsome bool
	//
	//fmt.Println("请输入身高(cm)")
	//fmt.Scanln(&height)
	//fmt.Println("请输入财富(千万)")
	//fmt.Scanln(&money)
	//fmt.Println("请输入是否帅(true/false)")
	//fmt.Scanln(&handsome)
	//
	//if height > 180 && money>1.0 && handsome{
	//	fmt.Println("我一定要嫁给他")
	//}else if height > 180 || money>1.0 || handsome{
	//	fmt.Println("嫁了吧,比上不足，比下有余")
	//}else{
	//	fmt.Println("不嫁！")
	//}

	//switch条件的比较操作，无需使用break跳出switch语句
	var key int
	fmt.Println("请输入一个字符:a,b,c,d,e,f,g")
	//fmt.Scanf("%d",&key)
	//go语言中一个读取字符的小问题
	fmt.Scanln(&key) //这里不能直接识别输入的字母，只能识别输入的数字
	//fmt.Scanln("%c",&key)
	fmt.Println(key)
	switch key{
	case 'a':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")
	case 'd':
		fmt.Println("周四")
	default:
		fmt.Println("输入有错误!")
	}
}
