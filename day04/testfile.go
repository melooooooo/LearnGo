package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	//编写一个程序，将一个文件的内容写入到另外一个文件中。这两个文件是已经存在的(使用ioutil.ReadFile和ioutil.WriteFile来实现
	file1path:="/Users/melooo/Desktop/test.txt"
	file2path:="/Users/melooo/Desktop/test1.txt"
	readFile, err := ioutil.ReadFile(file1path)
	if err!=nil{
		fmt.Println("read file error",err)
	}
	err = ioutil.WriteFile(file2path, readFile, 0666)
	if err!=nil{
		fmt.Println("write file error",err)
	}
}
