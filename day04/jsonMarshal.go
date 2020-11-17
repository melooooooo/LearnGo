package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	testStruct()
	fmt.Println()
	testFloat64()
	fmt.Println()
	testMap()
	fmt.Println()
	testslice()
}


//对于结构体的序列化，如果我们希望序列化之后的key的名字由我们自己指定，那么可以给struct指定一个tag标签
type Monster struct{
	Name string `json:"monster_name"`  //``里面包含的是tag标签
	Age int `json:"monster_age"`
	BrithDay string
	Sal float64
	Skill string
}

func testStruct(){
	monster:=Monster{
		Name:"牛魔王",
		Age:500,
		BrithDay: "2011-11-11",
		Sal: 8000.00,
		Skill: "牛魔拳",
	}
	data, err := json.Marshal(&monster)
	if err!=nil{
		fmt.Println("序列化错误",err)
	}
	fmt.Println("monster序列化之后为",string(data))
}


func testMap(){
	var a map[string]interface{}
	//使用map,需要make
	a=make(map[string]interface{})
	a["name"]="红孩儿"
	a["age"]=30
	a["address"]="洪崖洞"
	data, err := json.Marshal(a)
	if err!=nil{
		fmt.Println("序列化错误",err)
	}
	fmt.Printf("a map序列化之后的=%v",string(data))
}

func testslice(){
	var sli []map[string]interface{}
	var m1 map[string]interface{}
	m1=make(map[string]interface{})
	m1["name"]="jack"
	m1["age"]="7"
	m1["address"]="北京"
	sli = append(sli, m1)

	var m2 map[string]interface{}
	m2=make(map[string]interface{})
	m2["name"]="tom"
	m2["age"]=20
	m2["address"]=[2]string{"墨西哥","夏威夷"}
	sli = append(sli, m2)

	data, err := json.Marshal(sli)
	if err!=nil{
		fmt.Println("Marshal err",err)
	}

	fmt.Println("sli序列化之后",string(data))

}


func testFloat64(){
	var num1 float64=123.45
	data, err := json.Marshal(num1)
	if err!=nil{
		fmt.Println("Marshal error",err)
	}
	fmt.Printf("序列化之后的=%v",string(data))
}