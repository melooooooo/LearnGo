package main


//一个struct类型的数据类型中含有多个相同类型的基本数据类型
import "fmt"

type monster struct{
	Name string
	Age int
}

type E struct{
	monster
	int
	n int
}
func main(){
	var e E
	e.Age=300
	e.Name="狐狸精"
	e.int=20
	e.n=40
	fmt.Println("e=",e)
}
