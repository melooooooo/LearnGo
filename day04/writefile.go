package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	//O_CREATE如果不存在将会创建一个新文件
	file, err := os.OpenFile("/Users/melooo/Desktop/test.txt", os.O_WRONLY|os.O_CREATE, 0666)
	//os.OpenFile("")
	if err!=nil{
		fmt.Println("open file error",err)
		return
	}
	defer file.Close()
	str:="hello world!"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
