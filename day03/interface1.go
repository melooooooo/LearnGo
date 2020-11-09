package main

import "fmt"

/*interface可以定义一组方法，但是这些都不需要实现，并且interface不能包含任何变量
1.接口类的所有方法都没有方法体。即接口的方法奋斗史没有实现的方法。接口体现了程序设计的堕胎和高内聚迪偶尔的西乡。
2.golang中的接口，不需要显式的实现，只要一个变量，含有接口类型中所有的方法,那么这个变量就实现了这个接口。
因此，golang中没有interface这样的关键字

使用接口时需要注意的几点问题：
1.如果没有对interface进行初始化那么其默认将会是nil
2.空接口默认没有任何方法，所以所有类型都实现了空接口，我们可以把任何类型的变量赋给空接口
 */

type AInterface interface {
	Say()
}

type BInterface interface {
	Hello()
}

type Monster struct{

}

func (m Monster) Hello(){
	fmt.Println("monster hello~")
}
func (m Monster) Say(){
	fmt.Println("monster say~")
}
func main(){
	var monster Monster
	var a2 AInterface=monster
	var b2 BInterface=monster
	a2.Say()
	b2.Hello()
}

