package main

import "testing"


func TestMonster_Store(t *testing.T) {
	monster:=&Monster{
		//在结构体初始化的最后一行应该加上逗号,
		"红孩儿",10,"吐火",
	}
	res := monster.Store()
	if !res{
		t.Fatalf("期望为%v,实际为%v",true,res)
	}
	t.Logf("monster.Store()测试成功!")
}

func TestMonster_ReStore(t *testing.T) {
	var res =&Monster{}
	result := res.ReStore()
	if !result{
		t.Fatalf("期望为%v,实际为%v",true,result)
	}

	t.Logf("monster ReStore()测试成功!")
}
