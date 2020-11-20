package main

import "testing"

func TestCal(t *testing.T){
	res:=Add(1,2)
	if res!=3{
		t.Fatalf("add 的期望值为3,实际值为%v",res)
	}else{
		t.Logf("test Add success!")
	}
}
