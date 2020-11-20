package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct{
	Name string
	Age int
	Skill string
}

//给Monster绑定一个Store方法，可以将一个Monster变量(对象)序列化后保存到文件里

func ( monster *Monster) Store() bool{
	//先序列化
	data, err := json.Marshal(monster)
	if err!=nil{
		fmt.Println("marshal err=",err)
		return false
	}
	filePath:="/tmp/test.txt"
	err = ioutil.WriteFile(filePath,data, 0666)
	if err!=nil{
		fmt.Println("write file err=",err)
		return false
	}
	return true
}

func (monster *Monster) ReStore() bool{
	filePath:="/tmp/test.txt"
	file, err := ioutil.ReadFile(filePath)
	if err!=nil{
		fmt.Println("ReadFile err",err)
		return false
	}
	err = json.Unmarshal(file,monster)
	if err!=nil {
		fmt.Println("UnMarshal err",err)
		return false
	}
	return true

}