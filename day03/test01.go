package main

import "fmt"

type Cat struct{
	Name string
	Age int
	Color string
	Hobby string
}

func main(){
	var cat Cat
	cat.Name="小白"
	cat.Age=3
	cat.Color="白色"
	cat.Hobby="eat"
	fmt.Println(cat)

}
