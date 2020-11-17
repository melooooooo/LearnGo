package main

import (
	"fmt"
	"io/ioutil"
)

//使用io/ioutil来将文件一次性的读取结束

func main(){
	//_, err1 := ioutil.ReadFile("/Users/melooo/Desktop/test.txt")

	file, err := ioutil.ReadFile("/Users/melooo/Desktop/test.txt")
	//返回结果的类型是个byte数组
	if err!=nil{
		fmt.Println("read file err",err)
	}
	fmt.Printf("%v",string(file))
	//我们没有显式open文件，因此也不需要进行关闭文件的操作
}
