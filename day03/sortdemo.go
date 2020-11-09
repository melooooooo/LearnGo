package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct{
	Name string
	Age int
}

type HeroSlice []Hero

func (hs HeroSlice) Len() int{
	return len(hs)
}

func (hs HeroSlice) Less(i,j int) bool{
	return hs[i].Age<hs[j].Age
}

func (hs HeroSlice) Swap(i,j int){
	hs[i],hs[j]=hs[j],hs[i]
}
func main(){
	var intslice =[]int{0,-1,7,90,10}
	//可以直接使用系统提供的方法进行排序
	sort.Ints(intslice)
	fmt.Println(intslice)

	var hs HeroSlice
	for i := 0; i < 10; i++ {
		hero:=Hero{
			Name:fmt.Sprintf("英雄%02d",rand.Intn(100)),
			Age:rand.Intn(100),
		}
		//append(hs,hero)
		hs=append(hs,hero)
	}

	fmt.Println("sort前")
	for _,v := range hs{
		fmt.Println(v)
	}

	//调用sort.Sort
	sort.Sort(hs)


	fmt.Println("sort后")
	for _,v := range hs{
		fmt.Println(v)
	}

	//fmt.Println("sort后")
}
