package main

import (
	"encoding/json"
	"fmt"
)

type monster struct{
	Name string `json:"monster_name"`  //``里面包含的是tag标签
	Age int `json:"monster_age"`
	BrithDay string
	Sal float64
	Skill string
}


func main(){

	str:=" {\"monster_name\":\"牛魔王\",\"monster_age\":500,\"BrithDay\":\"2011-11-11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"
	var mon monster
	err := json.Unmarshal([]byte(str), &mon)
	if err!=nil{
		fmt.Println("unmarshal err=%v\n",err)
	}
	fmt.Printf("反序列化后 monster=%v monster.Name=%v\n",mon,mon.Name)
}


func unsharalSlice(){
	str:=" [{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"},{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":20,\"name\":\"tom\"}]"
	var sli []map[string]interface{}
	//反序列化，不需要make，因为make操作被封装在了Unmarshal函数当中
	err := json.Unmarshal([]byte(str), &sli)
	if err!=nil{
		fmt.Println("unmarshal err",err)
	}
	fmt.Println("反序列化 slice=", sli)
}
