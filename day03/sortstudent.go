package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct{
	Name string
	Age int
	Score float64
}

type Students []Student

func (s Students) Len() int{
	return len(s)
}

func (s Students) Less(i,j int) bool{
	return s[i].Score>s[j].Score
}

func (s Students) Swap(i,j int){
	s[i],s[j]=s[j],s[i]
}

func main(){
	var s Students
	for i := 0; i < 10; i++ {
		stu:=Student{
			Name:fmt.Sprintf("stu%02d",rand.Intn(100)),
			Age: rand.Intn(100),
			Score: rand.Float64()*100,
		}
		s = append(s,stu)
	}
	fmt.Println("排序前...")
	for _,v := range s{
		fmt.Println(v)
	}

	fmt.Println("排序之后...")

	sort.Sort(s)
	for _,v :=range s{
		fmt.Println(v)
	}

}
