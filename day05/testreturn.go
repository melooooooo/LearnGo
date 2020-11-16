package main

import (
	"fmt"
	"os"
)

func CopyFile(src,des string)(w int64,err error){
	srcFile,err:=os.Open(src)
	if err!=nil{
		return
	}
	defer srcFile.Close()
	return 888,nil
}

func main(){
	file, err := CopyFile("hahhah", "heiheieh")
	fmt.Println(file)
	fmt.Println(err)
}
