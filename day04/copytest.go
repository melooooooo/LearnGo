package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(desFilename string,srcFilename string)(written int64,err error){
	srcFile, err := os.Open(srcFilename)
	if err!=nil{
		fmt.Println("open file error",err)
		return 
	}
	reader := bufio.NewReader(srcFile)
	defer srcFile.Close()
	desFile, err := os.OpenFile(desFilename, os.O_WRONLY|os.O_CREATE, 0666)
	if err!=nil{
		fmt.Println("open file error",err)
		return 
	}
	writer := bufio.NewWriter(desFile)
	defer srcFile.Close()
	return io.Copy(writer,reader)
}
func main(){
	srcFile:="/Users/melooo/Desktop/test.txt"
	desFile:="/Users/melooo/Desktop/test2.txt"
	_, err := CopyFile(desFile, srcFile)
	if err!=nil {
		fmt.Println("copy file error", err)
	}else{
		fmt.Println("拷贝成功！")
	}
}