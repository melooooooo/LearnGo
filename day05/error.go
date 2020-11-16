package main

import (
	"errors"
	"fmt"
)

type myError struct{
	arg int
	errMsg string
}

func (e *myError) Error() string{
	return fmt.Sprintf("%d-%s",e.arg,e.errMsg)
}

func err_test(arg int) (int,error){
	if arg<0{
		return -1,errors.New("Bad arguments-negtive")
	}else if arg>256{
		return -1,&myError{arg,"Bad arguments-too large"}
	}
	return arg*arg,nil
}
func main(){
	for _,v := range []int{-1,4,1000}{
		if r,e:=err_test(v);e!=nil{
			fmt.Println("failed",e)
		}else{
			fmt.Println("success",r)
		}
	}
}


