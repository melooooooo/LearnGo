
package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile1(dst, src string) (w int64, err error) {
srcFile, err := os.Open(src)
if err != nil {
return
}

defer srcFile.Close()

dstFile, err := os.Create(dst)
if err != nil {
return
}

defer dstFile.Close()

return io.Copy(dstFile, srcFile)
}


func test1() (result int) {
	defer func() {
		result = 12
	}()
	return 10
}

func main() {
	fmt.Println(test1())     // 12
}