package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//通过bufio.NewReader(file)来进行读取文件
func main(){
	file, err := os.Open("/Users/melooo/Desktop/test.txt")
	if err!=nil{
		fmt.Println("open file err",err)
	}
	defer file.Close()
	//const(
	//	defaultBufSize=4096
	//)
	reader := bufio.NewReader(file)
	for{
		str, err := reader.ReadString('\n')
		if err==io.EOF{
			break
		}
		fmt.Println(str)
	}
	fmt.Println("文件读取结束...")
}
