package main

import "fmt"


type Count int
func (count *Count) Increment(){
	*count++
}

func (count *Count) Decrement(){
	*count--
}

func (count *Count) IsZero() bool{
	return *count==0
}

type Part struct{
	stat string
	Count
}

func (part Part) IsZero() bool{
	return part.Count.IsZero() && part.stat==""
}

func (part Part) String() string{
return fmt.Sprintf("<<%s,%d>>",part.stat,part.Count)
}
func main(){
	var i Count=-1
	fmt.Printf("Value of count before increment: %d\n", i)
	i.Increment()
	fmt.Printf("Value of count after increment: %d\n", i)
	part:=Part{"232",0}
	fmt.Printf("Part:%v\n",part)
	fmt.Printf("part is zero t/f?:%t\n",part.IsZero())
	fmt.Printf("Count in Part is zero t/f?: %t\n", part.Count.IsZero()) // 尽管覆盖了匿名字段的方法，单还是可以访问
}