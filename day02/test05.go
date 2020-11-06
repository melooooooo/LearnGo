package main

import "strings"

func main(){
	//str2:="hello北京"
	//r:=[]rune(str2)
	//for i := 0; i < len(r); i++ {
	//	fmt.Println("字符=%c\n",r[i])
	//	//fmt.Println("字符=%c\n",r[i])
	//}
	//str:="hello北"
	//res, err := strconv.Atoi("123")
	//if err!=nil{
	//	fmt.Println("转换错误")
	//}else{
	//	fmt.Println("转换的结果是",res)
	//}

	//Itoa和Atoi函数的作用是相反的效果
	//result := strconv.Itoa(123)
	//fmt.Printf("result type=%T,result=%v",result,result)


	//字符串转成byte[]
	//var bytes []byte =[]byte("hello go")
	//fmt.Printf("bytes=%v\n",bytes)

	//[]byte转字符串
	//str:=string([]byte{97,98,99})
	//fmt.Printf("str=%v",str)

	//10进制转成2，8，16进制
	//print(strconv.FormatInt(123, 2))

	//查找子串是否包含在指定的字符串当中
	//print(strings.Contains("seafood", "sea"))

	//统计一个字符串中有几个指定的子串
	//print(strings.Count("ceheese", "e"))

	//不区分大小写的比较
	//print(strings.EqualFold("Abc", "abc"))

	//返回第一次出现的下标
	//print(strings.Index("findtest", "test"))

	//返回最后一次出现的下标
	print(strings.LastIndex("findtesttest", "test"))

}
