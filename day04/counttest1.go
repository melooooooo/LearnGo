package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct{
	ChCount int
	NumCount int
	SpaceCount int
	OtherCount int
}
func main(){
	file, err := os.Open("/Users/melooo/Desktop/test.txt")
	if err!=nil{
		fmt.Println("open file error",err)
		return
	}
	reader := bufio.NewReader(file)
	defer file.Close()
	var count CharCount
	for{
		res, err1 := reader.ReadString('\n')
		if err1==io.EOF{
			break
		}
		str := []rune(res)
		for _,v:=range str{
			switch {
				case v>='a' && v<='z':
					fallthrough
			case v>='A' && v <='Z':
				count.ChCount++
			case v==' ' || v=='\t':
				count.SpaceCount++
			case v>='0'||v<='9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}

	}
	fmt.Printf("字符的个数为%v,数字的个数为%v,空格的个数为%v,其他字符的个数为%v",count.ChCount,count.NumCount,count.SpaceCount,count.OtherCount)
}
