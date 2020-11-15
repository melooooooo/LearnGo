package main

import "fmt"

type Usb interface{
	Start()
	Stop()
}

type Phone struct{
	name string
}

func (c Phone) Start(){
	fmt.Println("手机开始工作...")
}

func (c Phone) Stop(){
	fmt.Println("手机停止工作...")
}

func (c Phone) Call(){
	fmt.Println("手机正在打电话...")
}

type Camera struct{
	name string
}

func (c Camera) Start(){
	fmt.Println("相机开始工作...")
}

func (c Camera) Stop(){
	fmt.Println("相机停止工作...")
}

type Computer struct{

}

func (c Computer) Working(u Usb){
	u.Start()
	if phone,ok:=u.(Phone);ok{
		phone.Call()
	}
	u.Stop()
}


func main(){
	var usbArr [3]Usb
	usbArr[0]=Phone{"vivo"}
	usbArr[1]=Phone{"xiaomi"}
	usbArr[2]=Camera{"nikon"}

	//遍历usbArr
	var computer Computer
	//for i := 0; i < len(usbArr); i++ {
	//	computer.Working(usbArr[i])
	//	fmt.Println()
	//}
	for _,v
	fmt.Println(usbArr)
}