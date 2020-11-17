package main

import (
	"fmt"
	"os"
)

func main(){
	file, err := os.Open("/Users/melooo/Desktop/test.txt")
	//这里得到的file是一个指针 *File
	if err!=nil{
		fmt.Println("open file error",err)
	}
	fmt.Printf("file=%v",file)
	err = file.Close()
	if err!=nil{
		fmt.Println("close file error",err)
	}
}
