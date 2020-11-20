package main

import "fmt"

func main() {
	ch := make(chan string)
	// len(ch): channel 中剩余未读取的数据个数; cap(ch): channel 的容量
	fmt.Println("len(ch) = ", len(ch), "cap(ch) = ", cap(ch))
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println("i = ", i)
		}
		ch <- "子 go 程打印完毕"
	}()
	str := <-ch
	fmt.Println(str)
}

